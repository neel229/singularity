package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
	dataSource = "postgresql://root:postgres@localhost:5432/stockmarket-simulator?sslmode=disable"
)

// We will use testQueries throughout our application
// for testing the db queries
var testQueries *Queries
var testDB *sql.DB

// Entry test where we setup DB connection
// for testing
func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(driverName, dataSource)
	if err != nil {
		log.Fatalf("There was an error connecting the database: %v", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
