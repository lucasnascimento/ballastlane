package db

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestConnectToDB(t *testing.T) {
	// Load environment variables from a file
	err := godotenv.Load("../../.env.test")
	if err != nil {
		t.Fatalf("Error loading .env file")
	}

	// Call the function under test
	actualDb, err := ConnectToDB()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if actualDb == nil {
		t.Errorf("Expected a database connection, got nil")
	}

	// Close the database connection
	CloseDB(actualDb)
}
