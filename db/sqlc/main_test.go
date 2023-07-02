package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/focta/youtube-tech-school-golang/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M){

	// Lesson12 8 テストで使用しているconfigをutilのLaodConfig()での宣言文に置き換える
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(testDB)

	// m.Run() はテストが終了したかどうかを数値で返送してくれるので、その値をそのままos.Exitの終了値の指定に利用するというもの
	os.Exit(m.Run())
}