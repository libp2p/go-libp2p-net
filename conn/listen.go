package conn

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	ic "github.com/ipfs/go-libp2p/p2p/crypto"
	filter "github.com/ipfs/go-libp2p/p2p/net/filter"
	transport "github.com/ipfs/go-libp2p/p2p/net/transport"
	peer "github.com/ipfs/go-libp2p/p2p/peer"
	"gx/ipfs/QmQopLATEYMNg7dVqZRNDfeE2S1yKy8zrRh5xnYiuqeZBn/goprocess"
	goprocessctx "gx/ipfs/QmQopLATEYMNg7dVqZRNDfeE2S1yKy8zrRh5xnYiuqeZBn/goprocess/context"
	msmux "gx/ipfs/QmUeEcYJrzAEKdQXjzTxCgNZgc9sRuwharsvzzm5Gd2oGB/go-multistream"
	tec "gx/ipfs/QmWHgLqrghM9zw77nF6gdvT9ExQ2RB9pLxkd8sDHZf1rWb/go-temp-err-catcher"
	context "gx/ipfs/QmZy2y8t9zQH2a1b8q2ZSLKp17ATuJoCNxxyMFG5qFExpt/go-net/context"
	ma "gx/ipfs/QmcobAGsCjYt5DXoq9et9L8yR8er7o7Cu3DTvpaq12jYSz/go-multiaddr"
)

const SecioTag = "/secio/1.0.0"
const NoEncryptionTag = "/plaintext/1.0.0"

const connAcceptBuffer = 32
const NegotiateReadTimeout = time.Second * 20

var catcher = tec.TempErrCatcher{
	IsTemp: func(e error) bool {
		// ignore connection breakages up to this point. but log them
		if e == io.EOF {
			log.Debugf("listener ignoring conn with EOF: %s", e)
			return true
		}

		te, ok := e.(tec.Temporary)
		if ok {
			log.Debugf("listener ignoring conn with temporary err: %s", e)
			return te.Temporary()
		}
		return false
	},
}

// ConnWrapper is any function that wraps a raw multiaddr connection
type ConnWrapper func(transport.Conn) transport.Conn

// listener is an object that can accept connections. It implements Listener
type listener struct {
	transport.Listener

	local peer.ID    // LocalPeer is the identity of the local Peer
	privk ic.PrivKey // private key to use to initialize secure conns

	filters *filter.Filters

	wrapper ConnWrapper

	proc goprocess.Process

	mux *msmux.MultistreamMuxer

	incoming chan transport.Conn

	ctx context.Context
}

func (l *listener) teardown() error {
	defer log.Debugf("listener closed: %s %s", l.local, l.Multiaddr())
	return l.Listener.Close()
}

func (l *listener) Close() error {
	log.Debugf("listener closing: %s %s", l.local, l.Multiaddr())
	return l.proc.Close()
}

func (l *listener) String() string {
	return fmt.Sprintf("<Listener %s %s>", l.local, l.Multiaddr())
}

func (l *listener) SetAddrFilters(fs *filter.Filters) {
	l.filters = fs
}

// Accept waits for and returns the next connection to the listener.
// Note that unfortunately this
func (l *listener) Accept() (net.Conn, error) {
	for con := range l.incoming {
		c, err := newSingleConn(l.ctx, l.local, "", con)
		if err != nil {
			if catcher.IsTemporary(err) {
				continue
			}
			return nil, err
		}

		if l.privk == nil || EncryptConnections == false {
			log.Warning("listener %s listening INSECURELY!", l)
			return c, nil
		}
		sc, err := newSecureConn(l.ctx, l.privk, c)
		if err != nil {
			log.Infof("ignoring conn we failed to secure: %s %s", err, c)
			continue
		}
		return sc, nil
	}
	return nil, fmt.Errorf("listener is closed")
}

func (l *listener) Addr() net.Addr {
	return l.Listener.Addr()
}

// Multiaddr is the identity of the local Peer.
// If there is an error converting from net.Addr to ma.Multiaddr,
// the return value will be nil.
func (l *listener) Multiaddr() ma.Multiaddr {
	return l.Listener.Multiaddr()
}

// LocalPeer is the identity of the local Peer.
func (l *listener) LocalPeer() peer.ID {
	return l.local
}

func (l *listener) Loggable() map[string]interface{} {
	return map[string]interface{}{
		"listener": map[string]interface{}{
			"peer":    l.LocalPeer(),
			"address": l.Multiaddr(),
			"secure":  (l.privk != nil),
		},
	}
}

func (l *listener) handleIncoming() {
	var wg sync.WaitGroup
	defer func() {
		wg.Wait()
		close(l.incoming)
	}()

	for {
		maconn, err := l.Listener.Accept()
		if err != nil {
			if catcher.IsTemporary(err) {
				continue
			}
			log.Warningf("listener errored and will close: %s", err)
			return
		}

		log.Debugf("listener %s got connection: %s <---> %s", l, maconn.LocalMultiaddr(), maconn.RemoteMultiaddr())

		if l.filters != nil && l.filters.AddrBlocked(maconn.RemoteMultiaddr()) {
			log.Debugf("blocked connection from %s", maconn.RemoteMultiaddr())
			maconn.Close()
			continue
		}
		// If we have a wrapper func, wrap this conn
		if l.wrapper != nil {
			maconn = l.wrapper(maconn)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			maconn.SetReadDeadline(time.Now().Add(NegotiateReadTimeout))
			_, _, err = l.mux.Negotiate(maconn)
			if err != nil {
				log.Info("negotiation of crypto protocol failed: ", err)
				maconn.Close()
				return
			}

			// clear read readline
			maconn.SetReadDeadline(time.Time{})

			l.incoming <- maconn
		}()
	}
}

func WrapTransportListener(ctx context.Context, ml transport.Listener, local peer.ID, sk ic.PrivKey) (Listener, error) {
	l := &listener{
		Listener: ml,
		local:    local,
		privk:    sk,
		mux:      msmux.NewMultistreamMuxer(),
		incoming: make(chan transport.Conn, connAcceptBuffer),
		ctx:      ctx,
	}
	l.proc = goprocessctx.WithContextAndTeardown(ctx, l.teardown)

	if EncryptConnections {
		l.mux.AddHandler(SecioTag, nil)
	} else {
		l.mux.AddHandler(NoEncryptionTag, nil)
	}

	go l.handleIncoming()

	log.Debugf("Conn Listener on %s", l.Multiaddr())
	log.Event(ctx, "swarmListen", l)
	return l, nil
}

type ListenerConnWrapper interface {
	SetConnWrapper(ConnWrapper)
}

// SetConnWrapper assigns a maconn ConnWrapper to wrap all incoming
// connections with. MUST be set _before_ calling `Accept()`
func (l *listener) SetConnWrapper(cw ConnWrapper) {
	l.wrapper = cw
}
