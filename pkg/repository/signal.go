package repository

import (
	"database/sql"
	"time"
)

// SignalStore is an interface that defines the method for saving signal data.
// It abstracts the storage mechanism for signals.
type SignalStore interface {
	// Add saves a signal to the repository.
	Add(signal string, inserted_at time.Time) error
}

// SignalStoreImpl is an implementation of SignalStore that uses a SQL database.
type SignalStoreImpl struct {
	db *sql.DB
}

// NewSignalStore ...
func NewSignalStore(db *sql.DB) *SignalStoreImpl {
	return &SignalStoreImpl{db}
}

// Add implements the Add method of the SignalStore interface.
// It inserts a new signal record into the database.
// Parameters:
// - signal: The signal data to be saved.
// - inserted_at: The timestamp at which the signal was inserted.
// Returns:
// - error: An error if the operation fails, nil otherwise.
func (r *SignalStoreImpl) Add(signal string, inserted_at time.Time) error {
	_, err := r.db.Exec("INSERT INTO signals (signal, inserted_at) VALUES ($1, $2)", signal, inserted_at)
	return err
}
