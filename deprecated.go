package net

import (
	helpers "github.com/libp2p/go-libp2p/helpers"
	moved "github.com/libp2p/go-libp2p/skel/network"
)

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.MessageSizeMax instead.
const MessageSizeMax = moved.MessageSizeMax

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.Stream instead.
type Stream = moved.Stream

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.Direction instead.
type Direction = moved.Direction

const (
	// Deprecated: use github.com/libp2p/go-libp2p/skel/network.DirectionUnknown instead.
	DirUnknown = moved.DirUnknown
	// Deprecated: use github.com/libp2p/go-libp2p/skel/network.DirInbound instead.
	DirInbound = moved.DirInbound
	// Deprecated: use github.com/libp2p/go-libp2p/skel/network.DirOutbound instead.
	DirOutbound = moved.DirOutbound
)

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.Stat instead.
type Stat = moved.Stat

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.StreamHandler instead.
type StreamHandler = moved.StreamHandler

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.ConnSecurity instead.
type ConnSecurity = moved.ConnSecurity

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.ConnMultiaddrs instead.
type ConnMultiaddrs = moved.ConnMultiaddrs

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.Conn instead.
type Conn = moved.Conn

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.ConnHandler instead.
type ConnHandler = moved.ConnHandler

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.Network instead.
type Network = moved.Network

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.ErrNoRemoteAddrs instead.
var ErrNoRemoteAddrs = moved.ErrNoRemoteAddrs

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.ErrNoConn instead.
var ErrNoConn = moved.ErrNoConn

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.Dialer instead.
type Dialer = moved.Dialer

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.Connectedness instead.
type Connectedness = moved.Connectedness

const (
	// Deprecated: use github.com/libp2p/go-libp2p/skel/network.NotConnected instead.
	NotConnected = moved.NotConnected

	// Deprecated: use github.com/libp2p/go-libp2p/skel/network.Connected instead.
	Connected = moved.Connected

	// Deprecated: use github.com/libp2p/go-libp2p/skel/network.CanConnect instead.
	CanConnect = moved.CanConnect

	// Deprecated: use github.com/libp2p/go-libp2p/skel/network.CannotConnect instead.
	CannotConnect = moved.CannotConnect
)

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.Notifiee instead.
type Notifiee = moved.Notifiee

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.NotifyBundle instead.
type NotifyBundle = moved.NotifyBundle

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.WithNoDial instead.
var WithNoDial = moved.WithNoDial

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.GetNoDial instead.
var GetNoDial = moved.GetNoDial

// Deprecated: use github.com/libp2p/go-libp2p/helpers.EOFTimeout instead.
var EOFTimeout = helpers.EOFTimeout

// Deprecated: use github.com/libp2p/go-libp2p/helpers.ErrExpectedEOF instead.
var ErrExpectedEOF = helpers.ErrExpectedEOF

// Deprecated: use github.com/libp2p/go-libp2p/helpers.FullClose instead.
var FullClose = helpers.FullClose

// Deprecated: use github.com/libp2p/go-libp2p/helpers.AwaitEOF instead.
var AwaitEOF = helpers.AwaitEOF

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.DialPeerTimeout instead.
var DialPeerTimeout = moved.DialPeerTimeout

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.GetDialPeerTimeout instead.
var GetDialPeerTimeout = moved.GetDialPeerTimeout

// Deprecated: use github.com/libp2p/go-libp2p/skel/network.WithDialPeerTimeout instead.
var WithDialPeerTimeout = moved.WithDialPeerTimeout
