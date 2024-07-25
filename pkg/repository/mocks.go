package repository

import (
	"time"

	"github.com/stretchr/testify/mock"
)

// MockSignalStore is a mock type for the SignalStore interface.
// It embeds the mock.Mock struct from the testify package to provide mocking capabilities.
type MockSignalStore struct {
	mock.Mock
}

// Add is a mocked method that saves a signal to the repository.
func (m *MockSignalStore) Add(signal string, inserted_at time.Time) error {
	args := m.Called(signal, inserted_at)
	return args.Error(0)
}
