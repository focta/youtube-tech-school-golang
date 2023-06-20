package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	driver = "postgres"
	dbSource = "postgresql://yout:youtpass@localhost:15434/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M){
	conn, err := sql.Open(driver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(conn)

	// m.Run() はテストが終了したかどうかを数値で返送してくれるので、その値をそのままos.Exitの終了値の指定に利用するというもの
	os.Exit(m.Run())
}