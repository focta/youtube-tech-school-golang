package api

import (
	"net/http"

	db "github.com/focta/youtube-tech-school-golang/db/sqlc"
	"github.com/gin-gonic/gin"
)

// / Lesson11 6　リクエストで受け取る値を処理するための構造体を追記する(これはsqlcですでに生成したパラメータと同一のものがあるためコピペ)
type createAccountRequest struct {
	// Lesson11 7 バリデーションを追加する
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

// Lesson11 5　server.go のrouterで仮指定していた関数を生成する
func (server *Server) createAccount(ctx *gin.Context) {
	// Lesson11 8　リクエストの構造体とリクエストをバインド処理する
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// Lesson11 9　失敗時のJsonレスポンスの処理を追加する(ただしerrorResponse()は自前の関数で、まだ実装していないので、この時点ではエラーとなる。)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Lesson11 12　DB処理を追記していく。
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
// Lesson11 13　正常系のレスポンスを記載する。
	ctx.JSON(http.StatusOK, account)
}
