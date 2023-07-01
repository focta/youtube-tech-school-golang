package api

import (
	"database/sql"
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

// Lesson11 31 忘れていたのでリクエストのマッピング用の構造体を宣言しておく。
type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// Lesson11 25　serverで作成したルーティングの関数を作成する。
func (server *Server) getAccount(ctx *gin.Context) {
	// Lesson11 26 ShouldBindUri()でパラメータのbindを行う(この時点では　getAccountRequest型を宣言していないため、エラーとなっている)
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Lesson11 27 アカウントをDBから引っ張ってくる
	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		// Lesson11 28 アカウントがテーブルに無い場合のエラーを定義する
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		// Lesson11 29 上記以外のデータはSQL発行での構成上のエラーとして500で返送する
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// Lesson11 30 正常系のレスポンスを設定する
	ctx.JSON(http.StatusOK, account)

}

// Lesson11 33 リストでアカウントを取得するための構造体を実装する
type getListAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// Lesson11 34　listで取得するメソッドをgetAccountをkぴーして実装する
func (server *Server) listAccount(ctx *gin.Context) {
	var req getListAccountRequest
	// Lesson11 36　 今回はURLパラメータから取得するので、Queryの取得メソッドを呼び出すように変更
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Lesson11 35　ListAccountのパラメータを設定して、メソッドの引数に指定する
	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// Lesson11 30 正常系のレスポンスを設定する
	ctx.JSON(http.StatusOK, accounts)

}
