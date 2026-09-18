package backoff

import (
	"context"
	"math"
	"math/rand/v2"
	"time"
)

const defaultBaseWait = 10 * time.Millisecond

type strategy uint8

const (
	quadratic strategy = iota
	exponential
)

// Backoff implements quadratic or exponential backoff with jitter.
// It must not be used by multiple goroutines concurrently.
type Backoff struct {
	n        uint64 // consecutive failures
	baseWait time.Duration
	maxWait  time.Duration
	strategy strategy
}

// NewBackoff creates a quadratic backoff with delays of baseWait*n*n for
// consecutive failures n starting at 1, jittered by +/-20% and capped at maxWait.
// baseWait is floored at 10ms; maxWait is floored at the resulting baseWait.
func NewBackoff(baseWait, maxWait time.Duration) *Backoff {
	baseWait = max(baseWait, defaultBaseWait)
	maxWait = max(maxWait, baseWait)
	return &Backoff{
		baseWait: baseWait,
		maxWait:  maxWait,
	}
}

// NewExponentialBackoff creates a backoff that doubles the base wait with each
// failure, starting at baseWait.
// baseWait is floored at 10ms; maxWait is floored at the resulting baseWait.
func NewExponentialBackoff(baseWait, maxWait time.Duration) *Backoff {
	b := NewBackoff(baseWait, maxWait)
	b.strategy = exponential
	return b
}

// Wait pauses for the next backoff interval or returns early if the context is done.
func (b *Backoff) Wait(ctx context.Context) {
	b.WaitOn(ctx, nil)
}

// WaitOn pauses for the next backoff interval, returning early if ctx is done or
// a value is received on wake. An early wake still advances the interval; an
// already-cancelled ctx does not.
func (b *Backoff) WaitOn(ctx context.Context, wake <-chan struct{}) {
	if ctx.Err() != nil {
		return
	}

	wait := b.nextDelay(rand.Float64())
	select {
	case <-ctx.Done():
	case <-wake:
	case <-time.After(wait):
	}
}

// nextDelay counts a failure and returns the interval to wait for it, jittered
// by random (expected in [0, 1]), capped at maxWait.
func (b *Backoff) nextDelay(random float64) time.Duration {
	b.n++
	jitter := 0.8 + 0.4*random
	wait := float64(b.rawDelay()) * jitter
	if wait >= float64(b.maxWait) {
		return b.maxWait
	}
	return time.Duration(wait)
}

// rawDelay saturates above the largest duration even after minimum jitter.
func (b *Backoff) rawDelay() uint64 {
	base := uint64(b.baseWait)
	if b.strategy == exponential {
		if base > math.MaxUint64>>(b.n-1) {
			return math.MaxUint64
		}
		return base << (b.n - 1)
	}
	if base > math.MaxUint64/b.n/b.n {
		return math.MaxUint64
	}
	return base * b.n * b.n
}

// Reset resets the backoff counter.
func (b *Backoff) Reset() {
	b.n = 0
}
