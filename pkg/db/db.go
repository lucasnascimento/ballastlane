package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// ConnectToDB attempts to verify the database connection by calling the Ping method on the *sql.DB object.
// If the Ping call is successful, it prints "Successfully connected!" to the standard output and returns the *sql.DB object.
// If there is an error at any point (either from the initial database setup or the Ping call), it returns nil and the error.
//
// Returns:
// - *sql.DB: A pointer to the sql.DB object representing the database connection. This is returned if the connection is successfully verified.
// - error: An error object that indicates why the connection verification failed. This is nil if the connection is successful.
func ConnectToDB() (*sql.DB, error) {
	// Usando variáveis de ambiente para a string de conexão
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Construindo a string de conexão
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Conectando ao banco de dados
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// Verificando a conexão
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")
	return db, nil
}

// CloseDB attempts to close the database connection represented by the *sql.DB object passed as an argument.
// It calls the Close method on the *sql.DB object.
// If the Close method returns an error, CloseDB logs the error using log.Fatal, which also terminates the program.
// This function does not return any value.
//
// Parameters:
// - db: A pointer to the sql.DB object that represents the database connection to be closed.
func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
