package net

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/helpers"
	moved "github.com/libp2p/go-libp2p-core/network"
)

// Deprecated: use github.com/libp2p/go-libp2p-core/network.MessageSizeMax instead.
const MessageSizeMax = moved.MessageSizeMax

// Deprecated: use github.com/libp2p/go-libp2p-core/network.Stream instead.
type Stream = moved.Stream

// Deprecated: use github.com/libp2p/go-libp2p-core/network.Direction instead.
type Direction = moved.Direction

const (
	// Deprecated: use github.com/libp2p/go-libp2p-core/network.DirectionUnknown instead.
	DirUnknown = moved.DirUnknown
	// Deprecated: use github.com/libp2p/go-libp2p-core/network.DirInbound instead.
	DirInbound = moved.DirInbound
	// Deprecated: use github.com/libp2p/go-libp2p-core/network.DirOutbound instead.
	DirOutbound = moved.DirOutbound
)

// Deprecated: use github.com/libp2p/go-libp2p-core/network.Stat instead.
type Stat = moved.Stat

// Deprecated: use github.com/libp2p/go-libp2p-core/network.StreamHandler instead.
type StreamHandler = moved.StreamHandler

// Deprecated: use github.com/libp2p/go-libp2p-core/network.ConnSecurity instead.
type ConnSecurity = moved.ConnSecurity

// Deprecated: use github.com/libp2p/go-libp2p-core/network.ConnMultiaddrs instead.
type ConnMultiaddrs = moved.ConnMultiaddrs

// Deprecated: use github.com/libp2p/go-libp2p-core/network.Conn instead.
type Conn = moved.Conn

// Deprecated: use github.com/libp2p/go-libp2p-core/network.ConnHandler instead.
type ConnHandler = moved.ConnHandler

// Deprecated: use github.com/libp2p/go-libp2p-core/network.Network instead.
type Network = moved.Network

// Deprecated: use github.com/libp2p/go-libp2p-core/network.ErrNoRemoteAddrs instead.
var ErrNoRemoteAddrs = moved.ErrNoRemoteAddrs

// Deprecated: use github.com/libp2p/go-libp2p-core/network.ErrNoConn instead.
var ErrNoConn = moved.ErrNoConn

// Deprecated: use github.com/libp2p/go-libp2p-core/network.Dialer instead.
type Dialer = moved.Dialer

// Deprecated: use github.com/libp2p/go-libp2p-core/network.Connectedness instead.
type Connectedness = moved.Connectedness

const (
	// Deprecated: use github.com/libp2p/go-libp2p-core/network.NotConnected instead.
	NotConnected = moved.NotConnected

	// Deprecated: use github.com/libp2p/go-libp2p-core/network.Connected instead.
	Connected = moved.Connected

	// Deprecated: use github.com/libp2p/go-libp2p-core/network.CanConnect instead.
	CanConnect = moved.CanConnect

	// Deprecated: use github.com/libp2p/go-libp2p-core/network.CannotConnect instead.
	CannotConnect = moved.CannotConnect
)

// Deprecated: use github.com/libp2p/go-libp2p-core/network.Notifiee instead.
type Notifiee = moved.Notifiee

// Deprecated: use github.com/libp2p/go-libp2p-core/network.NotifyBundle instead.
type NotifyBundle = moved.NotifyBundle

// Deprecated: use github.com/libp2p/go-libp2p-core/network.WithNoDial instead.
func WithNoDial(ctx context.Context, reason string) context.Context {
	return moved.WithNoDial(ctx, reason)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/network.GetNoDial instead.
func GetNoDial(ctx context.Context) (nodial bool, reason string) {
	return moved.GetNoDial(ctx)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/helpers.EOFTimeout instead.
var EOFTimeout = helpers.EOFTimeout

// Deprecated: use github.com/libp2p/go-libp2p-core/helpers.ErrExpectedEOF instead.
var ErrExpectedEOF = helpers.ErrExpectedEOF

// Deprecated: use github.com/libp2p/go-libp2p-core/helpers.FullClose instead.
func FullClose(s moved.Stream) error {
	return helpers.FullClose(s)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/helpers.AwaitEOF instead.
func AwaitEOF(s moved.Stream) error {
	return helpers.AwaitEOF(s)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/network.DialPeerTimeout instead.
// Warning: this variable's type makes it impossible to alias by reference. Reads and writes
// from/to this variable may be inaccurate or not have the intended effect.
var DialPeerTimeout = moved.DialPeerTimeout

// Deprecated: use github.com/libp2p/go-libp2p-core/network.GetDialPeerTimeout instead.
func GetDialPeerTimeout(ctx context.Context) time.Duration {
	return moved.GetDialPeerTimeout(ctx)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/network.WithDialPeerTimeout instead.
func WithDialPeerTimeout(ctx context.Context, timeout time.Duration) context.Context {
	return moved.WithDialPeerTimeout(ctx, timeout)
}
