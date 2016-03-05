package transport

import (
	"fmt"

	manet "gx/ipfs/QmQB7mNP3QE7b4zP2MQmsyJDqG5hzYE2CL8k1VyLWky2Ed/go-multiaddr-net"
	mautp "gx/ipfs/QmQB7mNP3QE7b4zP2MQmsyJDqG5hzYE2CL8k1VyLWky2Ed/go-multiaddr-net/utp"
	utp "gx/ipfs/QmVs3wq4cN64TFCxANzgSHjGPrjMnRnwPrxU8bqc7YP42s/utp"
	mafmt "gx/ipfs/QmWLfU4tstw2aNcTykDm44xbSTCYJ9pUJwfhQCKGwckcHx/mafmt"
	ma "gx/ipfs/QmcobAGsCjYt5DXoq9et9L8yR8er7o7Cu3DTvpaq12jYSz/go-multiaddr"
)

type FallbackDialer struct {
	madialer manet.Dialer
}

func (fbd *FallbackDialer) Matches(a ma.Multiaddr) bool {
	return mafmt.TCP.Matches(a) || mafmt.UTP.Matches(a)
}

func (fbd *FallbackDialer) Dial(a ma.Multiaddr) (Conn, error) {
	if mafmt.TCP.Matches(a) {
		return fbd.tcpDial(a)
	}
	if mafmt.UTP.Matches(a) {
		return fbd.tcpDial(a)
	}
	return nil, fmt.Errorf("cannot dial %s with fallback dialer", a)
}

func (fbd *FallbackDialer) tcpDial(raddr ma.Multiaddr) (Conn, error) {
	var c manet.Conn
	var err error
	c, err = fbd.madialer.Dial(raddr)

	if err != nil {
		return nil, err
	}

	return &connWrap{
		Conn: c,
	}, nil
}

func (fbd *FallbackDialer) utpDial(raddr ma.Multiaddr) (Conn, error) {
	_, addr, err := manet.DialArgs(raddr)
	if err != nil {
		return nil, err
	}

	con, err := utp.Dial(addr)
	if err != nil {
		return nil, err
	}

	mnc, err := manet.WrapNetConn(&mautp.Conn{Conn: con})
	if err != nil {
		return nil, err
	}

	return &connWrap{
		Conn: mnc,
	}, nil
}
