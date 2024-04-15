package retry

import (
	"sync"

	"github.com/cenkalti/backoff/v4"
)

func Do(f func() error) error {
	b := get()
	defer put(b)
	return backoff.Retry(f, b)
}

func get() *backoff.ExponentialBackOff {
	b := pool.Get().(*backoff.ExponentialBackOff)
	b.Reset()
	return b
}

func put(b *backoff.ExponentialBackOff) {
	pool.Put(reset(b))
}

func reset(b *backoff.ExponentialBackOff) *backoff.ExponentialBackOff {
	*b = backoff.ExponentialBackOff{
		InitialInterval:     backoff.DefaultInitialInterval,
		RandomizationFactor: backoff.DefaultRandomizationFactor,
		Multiplier:          backoff.DefaultMultiplier,
		MaxInterval:         backoff.DefaultMaxInterval,
		MaxElapsedTime:      backoff.DefaultMaxElapsedTime,
		Stop:                backoff.Stop,
		Clock:               backoff.SystemClock,
	}
	return b
}

var pool = &sync.Pool{New: func() any { return backoff.NewExponentialBackOff() }}
