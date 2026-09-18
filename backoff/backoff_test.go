package backoff

import (
	"context"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var backoffConstructors = map[string]func(time.Duration, time.Duration) *Backoff{
	"quadratic":   NewBackoff,
	"exponential": NewExponentialBackoff,
}

func TestBackoffNextDelay(t *testing.T) {
	backoff := NewBackoff(time.Second, time.Minute)
	for _, expected := range []time.Duration{
		1 * time.Second,
		4 * time.Second,
		9 * time.Second,
		16 * time.Second,
		25 * time.Second,
		36 * time.Second,
		49 * time.Second,
		time.Minute,
		time.Minute,
	} {
		require.Equal(t, expected, backoff.nextDelay(0.5))
	}

	backoff.Reset()
	require.Equal(t, time.Second, backoff.nextDelay(0.5))
}

func TestBackoffNextDelayJitter(t *testing.T) {
	for name, newBackoff := range backoffConstructors {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, 800*time.Millisecond, newBackoff(time.Second, time.Minute).nextDelay(0))
			require.Equal(t, time.Second, newBackoff(time.Second, time.Minute).nextDelay(0.5))
			require.Equal(t, 1200*time.Millisecond, newBackoff(time.Second, time.Minute).nextDelay(1))

			require.Equal(t, time.Minute, newBackoff(time.Minute, time.Minute).nextDelay(1))

			backoff := newBackoff(time.Minute, time.Minute)
			require.Equal(t, 48*time.Second, backoff.nextDelay(0))
			require.Equal(t, time.Minute, backoff.nextDelay(0))
			require.Equal(t, time.Minute, backoff.nextDelay(0))
			require.Equal(t, uint64(3), backoff.n)
		})
	}
}

func TestBackoffDefaultBaseWait(t *testing.T) {
	for name, newBackoff := range backoffConstructors {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, defaultBaseWait, newBackoff(0, time.Minute).nextDelay(0.5))
			require.Equal(t, defaultBaseWait, newBackoff(-time.Second, time.Minute).nextDelay(0.5))
		})
	}
}

func TestBackoffWaitOnDoneContextDoesNotCountFailure(t *testing.T) {
	for name, newBackoff := range backoffConstructors {
		t.Run(name, func(t *testing.T) {
			backoff := newBackoff(time.Second, time.Minute)
			ctx, cancel := context.WithCancel(context.Background())
			cancel()

			backoff.Wait(ctx)
			require.Equal(t, time.Second, backoff.nextDelay(0.5))
		})
	}
}

func TestBackoffWaitOnReturnsOnWake(t *testing.T) {
	for name, newBackoff := range backoffConstructors {
		t.Run(name, func(t *testing.T) {
			backoff := newBackoff(time.Hour, time.Hour)
			wake := make(chan struct{}, 1)
			wake <- struct{}{}

			started := time.Now()
			backoff.WaitOn(context.Background(), wake)
			require.Less(t, time.Since(started), time.Second)
			require.Equal(t, uint64(1), backoff.n)
		})
	}
}

func TestExponentialBackoffNextDelay(t *testing.T) {
	backoff := NewExponentialBackoff(time.Second, time.Minute)
	for _, expected := range []time.Duration{
		time.Second,
		2 * time.Second,
		4 * time.Second,
		8 * time.Second,
		16 * time.Second,
		32 * time.Second,
		time.Minute,
		time.Minute,
		time.Minute,
	} {
		require.Equal(t, expected, backoff.nextDelay(0.5))
	}

	backoff.Reset()
	require.Equal(t, time.Second, backoff.nextDelay(0.5))
}

func TestBackoffDurationOverflow(t *testing.T) {
	for name, newBackoff := range backoffConstructors {
		t.Run(name, func(t *testing.T) {
			for _, random := range []float64{0, 0.5, 1} {
				backoff := newBackoff(time.Duration(math.MaxInt64/2), time.Duration(math.MaxInt64))
				require.Positive(t, backoff.nextDelay(random))
				require.Positive(t, backoff.nextDelay(random))
				for range 100 {
					require.Equal(t, time.Duration(math.MaxInt64), backoff.nextDelay(random))
				}
			}
			require.Equal(t, time.Duration(math.MaxInt64),
				newBackoff(time.Duration(math.MaxInt64), time.Duration(math.MaxInt64)).nextDelay(0.5))
		})
	}
}

func TestBackoffLargeFailureCount(t *testing.T) {
	for _, tc := range []struct {
		name string
		b    *Backoff
		n    uint64
	}{
		{"quadratic", NewBackoff(time.Nanosecond, time.Duration(math.MaxInt64)), 3_000_000_000},
		{"exponential", NewExponentialBackoff(time.Nanosecond, time.Duration(math.MaxInt64)), 63},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tc.b.n = tc.n
			require.Positive(t, tc.b.nextDelay(0))
			require.Equal(t, time.Duration(math.MaxInt64), tc.b.nextDelay(1))
		})
	}
}
