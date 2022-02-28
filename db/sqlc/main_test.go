package db

import (
	"database/sql"
	"log"
	"os"
	"simple_bank/db/util"
	"testing"

	_ "github.com/lib/pq"
)



var testQueries *Queries
var testDB *sql.DB
func TestMain(m *testing.M) {
	var err error 
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config file", err)
		return
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}