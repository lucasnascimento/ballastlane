package repository

import (
	"time"

	"github.com/stretchr/testify/mock"
)

// MockSignalsRepository is a mock type for the SignalsRepository interface.
// It embeds the mock.Mock struct from the testify package to provide mocking capabilities.
type MockSignalsRepository struct {
	mock.Mock
}

// SaveSignal is a mocked method that saves a signal to the repository.
func (m *MockSignalsRepository) SaveSignal(signal string, inserted_at time.Time) error {
	args := m.Called(signal, inserted_at)
	return args.Error(0)
}
