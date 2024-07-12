package repository

import (
	"database/sql"
	"time"
)

// ISignalsRepository is an interface that defines the method for saving signal data.
// It abstracts the storage mechanism for signals.
type ISignalsRepository interface {
	SaveSignal(signal string, inserted_at time.Time) error
}

// SignalsRepository is an implementation of ISignalsRepository that uses a SQL database.
type SignalsRepository struct {
	db *sql.DB
}

// NewSignalsRepository ...
func NewSignalsRepository(db *sql.DB) *SignalsRepository {
	return &SignalsRepository{db}
}

// SaveSignal implements the SaveSignal method of the ISignalsRepository interface.
// It inserts a new signal record into the database.
// Parameters:
// - signal: The signal data to be saved.
// - inserted_at: The timestamp at which the signal was inserted.
// Returns:
// - error: An error if the operation fails, nil otherwise.
func (r *SignalsRepository) SaveSignal(signal string, inserted_at time.Time) error {
	_, err := r.db.Exec("INSERT INTO signals (signal, inserted_at) VALUES ($1, $2)", signal, inserted_at)
	return err
}
