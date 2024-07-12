package clock

import "time"

// MockTimeProvider is a mock implementation of the TimeProvider interface.
type MockTimeProvider struct {
	FixedTime time.Time
}

// Now returns the fixed time set in the MockTimeProvider.
func (mtp MockTimeProvider) Now() time.Time {
	return mtp.FixedTime
}

// MockedExitProvider is a mock implementation of the ExitProvider interface.
type MockedExitProvider struct{}

// Exit does nothing because on tests we don't want to exit the program.
func (m MockedExitProvider) Exit(code int) {
	// Do nothing
}
