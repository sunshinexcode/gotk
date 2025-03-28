package vlimit

import (
	"time"

	"golang.org/x/time/rate"
)

// Every converts a minimum time interval between events to a Limit.
func Every(interval time.Duration) rate.Limit {
	return rate.Every(interval)
}

// New returns a new Limiter that allows events up to rate r and permits
// bursts of at most b tokens.
func New(r rate.Limit, b int) *rate.Limiter {
	return rate.NewLimiter(r, b)
}
