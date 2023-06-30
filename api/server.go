package api

import (
	db "github.com/focta/youtube-tech-school-golang/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Lesson11 ①構造体を作る
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// Lesson11 ②関数の外形を作る
func NewServer(store *db.Store) *Server {
	// Lesson11 3　server/routerまでの実装を行う
	server := &Server{store: store}
	router := gin.Default()

	// Lesson11 4　各パスを実装してみる(この時点では宣言した第2引数の関数は未実装のため、エラー状態)
	router.POST("/accounts", server.createAccount)
	// Lesson11 24　GETのパスを追加する
	router.GET("/accounts/:id", server.getAccount)

	server.router = router
	return server
}

// Lesson11 14　サーバーの起動するための関数を作成する
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// Lesson11 10　account.goでエラーとなっているerrorResponse関数を作成する
// Lesson11 11　gin.Hの返送の型を追加してやる
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
