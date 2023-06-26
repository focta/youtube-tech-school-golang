package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	driver = "postgres"
	dbSource = "postgresql://yout:youtpass@localhost:15434/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M){
	var err error
	testDB, err = sql.Open(driver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(testDB)

	// m.Run() はテストが終了したかどうかを数値で返送してくれるので、その値をそのままos.Exitの終了値の指定に利用するというもの
	os.Exit(m.Run())
}