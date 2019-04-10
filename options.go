package net

import (
	"context"
)

type noDialCtxKey struct{}

// NoDial is a context option that instructs the network to not attempt a new
// dial when opening a stream. The value of the key should be a string indicating
// the source of the option.
var NoDial = noDialCtxKey{}

func WithNoDial(ctx context.Context, src string) {
	context.WithValue(ctx, NoDial, src)
}
