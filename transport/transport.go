package transport

import (
	"net"
	"time"

	logging "QmWRypnfEwrgH4k93KEHN5hng7VjKYkWmzDYRuTZeh2Mgh/go-log"
	ma "QmaA6aDzeHjZiuqBtgYRz8ZXb1qMCoyMHgyDjBEYQniUKF/go-multiaddr"
	manet "QmanZCL6SXRfafiUEMCBLq2QR171uQSdXQ8YAdHXLd8Cwr/go-multiaddr-net"
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
