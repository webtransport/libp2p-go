package network

import (
	"context"
	"time"
)

// DialPeerTimeout is the default timeout for a single call to `DialPeer`. When
// there are multiple concurrent calls to `DialPeer`, this timeout will apply to
// each independently.
var DialPeerTimeout = 60 * time.Second

type noDialCtxKey struct{}
type dialPeerTimeoutCtxKey struct{}
type forceDirectDialCtxKey struct{}

var noDial = noDialCtxKey{}
var forceDirectDial = forceDirectDialCtxKey{}

// EXPERIMENTAL
// WithForceDirectDial constructs a new context with an option that instructs the network
// to attempt to force a direct connection to a peer via a dial even if a proxied connection to it already exists.
func WithForceDirectDial(ctx context.Context, reason string) context.Context {
	return context.WithValue(ctx, forceDirectDial, reason)
}

// EXPERIMENTAL
// GetForceDirectDial returns true if the force direct dial option is set in the context.
func GetForceDirectDial(ctx context.Context) (forceDirect bool, reason string) {
	v := ctx.Value(forceDirectDial)
	if v != nil {
		return true, v.(string)
	}

	return false, ""
}

// WithNoDial constructs a new context with an option that instructs the network
// to not attempt a new dial when opening a stream.
func WithNoDial(ctx context.Context, reason string) context.Context {
	return context.WithValue(ctx, noDial, reason)
}

// GetNoDial returns true if the no dial option is set in the context.
func GetNoDial(ctx context.Context) (nodial bool, reason string) {
	v := ctx.Value(noDial)
	if v != nil {
		return true, v.(string)
	}

	return false, ""
}

// GetDialPeerTimeout returns the current DialPeer timeout (or the default).
func GetDialPeerTimeout(ctx context.Context) time.Duration {
	if to, ok := ctx.Value(dialPeerTimeoutCtxKey{}).(time.Duration); ok {
		return to
	}
	return DialPeerTimeout
}

// WithDialPeerTimeout returns a new context with the DialPeer timeout applied.
//
// This timeout overrides the default DialPeerTimeout and applies per-dial
// independently.
func WithDialPeerTimeout(ctx context.Context, timeout time.Duration) context.Context {
	return context.WithValue(ctx, dialPeerTimeoutCtxKey{}, timeout)
}
