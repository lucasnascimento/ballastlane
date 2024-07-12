package clock

import "time"

// TimeProvider is an interface that defines the method for getting the current time.
type TimeProvider interface {
	Now() time.Time
}

// RealTimeProvider is an implementation of TimeProvider that returns the actual current time.
type RealTimeProvider struct{}

// Now returns the current time.
func (rtp RealTimeProvider) Now() time.Time {
	return time.Now()
}

