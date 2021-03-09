package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/neel229/forum/pkg/util"
)

// We will use testQueries throughout our application
// for testing the db queries
var testQueries *Queries
var testDB *sql.DB

// Entry test where we setup DB connection
// for testing
func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("There was an error connecting the database: %v", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
