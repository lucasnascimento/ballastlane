package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSaveSignal(t *testing.T) {
	db, mock, err := sqlmock.New() // Create a new mock SQL database
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewSignalsRepository(db) // Use the mock database in the repository

	// Expectation: there should be an exec action with the correct SQL query
	mock.ExpectExec("INSERT INTO signals \\(signal, inserted_at\\) VALUES \\(\\$1, \\$2\\)").WithArgs("test_signal", time.Time{}).WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the method to test
	err = repo.SaveSignal("test_signal", time.Time{})
	if err != nil {
		t.Errorf("error was not expected while saving signal: %s", err)
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
