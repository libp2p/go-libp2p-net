package transport

import (
	"net"
	"time"

	manet "gx/ipfs/QmQB7mNP3QE7b4zP2MQmsyJDqG5hzYE2CL8k1VyLWky2Ed/go-multiaddr-net"
	logging "gx/ipfs/Qmazh5oNUVsDZTs2g59rq8aYQqwpss8tcUWQzor5sCCEuH/go-log"
	ma "gx/ipfs/QmcobAGsCjYt5DXoq9et9L8yR8er7o7Cu3DTvpaq12jYSz/go-multiaddr"
)

var log = logging.Logger("transport")

type Conn interface {
	manet.Conn

	Transport() Transport
}

type Transport interface {
	Dialer(laddr ma.Multiaddr, opts ...DialOpt) (Dialer, error)
	Listen(laddr ma.Multiaddr) (Listener, error)
	Matches(ma.Multiaddr) bool
}

type Dialer interface {
	Dial(raddr ma.Multiaddr) (Conn, error)
	Matches(ma.Multiaddr) bool
}

type Listener interface {
	Accept() (Conn, error)
	Close() error
	Addr() net.Addr
	Multiaddr() ma.Multiaddr
}

type connWrap struct {
	manet.Conn
	transport Transport
}

func (cw *connWrap) Transport() Transport {
	return cw.transport
}

type DialOpt interface{}
type TimeoutOpt time.Duration
type ReuseportOpt bool

var ReusePorts ReuseportOpt = true
