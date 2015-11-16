package transport

import (
	"net"
	"time"

	manet "QmU5s159q8cZuM1f9Vqti4LHu6y8zyVc5dxv2py81sdp6Q/go-multiaddr-net"
	logging "QmWRypnfEwrgH4k93KEHN5hng7VjKYkWmzDYRuTZeh2Mgh/go-log"
	ma "QmbWxL1aXQhBjc1XGjGF1f2KGBMCBYSuT2ThA8YXnXJK83/go-multiaddr"
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

func IsTcpMultiaddr(a ma.Multiaddr) bool {
	p := a.Protocols()
	return len(p) == 2 && (p[0].Name == "ip4" || p[0].Name == "ip6") && p[1].Name == "tcp"
}

func IsUtpMultiaddr(a ma.Multiaddr) bool {
	p := a.Protocols()
	return len(p) == 3 && p[2].Name == "utp"
}
