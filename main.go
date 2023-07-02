package main

import (
	"database/sql"
	"log"

	// Lesson11 23　ドライバの実態をimportに追加する
	"github.com/focta/youtube-tech-school-golang/api"
	db "github.com/focta/youtube-tech-school-golang/db/sqlc"
	"github.com/focta/youtube-tech-school-golang/util"
	_ "github.com/lib/pq"
)

// Lesson11 16　db/sqlc/main_test.go からDB関連の接続情報をコピーする
const (
	driver   = "postgres"
	dbSource = "postgresql://yout:youtpass@localhost:15434/simple_bank?sslmode=disable"
	// Lesson11 21　serverAddressを定数として宣言
	serverAddress = "0.0.0.0:8080"
)

// Lesson11 15　メイン関数を作成する
func main() {

	// Lesson12 6 config.goの設定を追加する
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Lesson11 17　db/sqlc/main_test.go からDBの接続の処理を持ってくる。
	conn, err := sql.Open(config.DBDriver, config.DBSource) // Lesson12 7 config.goで読み込んだ設定を指定する
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	// Lesson11 18　dbのコネクションを開く
	store := db.NewStore(conn)
	// Lesson11 19　apiにdbのコネクション情報を渡す
	server := api.NewServer(store)

	// Lesson11 20　serverを開始する(apiパッケージ内のstartメソッドを利用する) ただし、この時点では 変数:serverAddress がまだ未定義なのでエラー
	err = server.Start(config.ServerAddress) // Lesson12 7 config.goで読み込んだ設定を指定する
	// Lesson11 22　Startメソッドでのエラーをチェックする
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
