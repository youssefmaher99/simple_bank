package api

import (
	"database/sql"
	"fmt"
	"net/http"
	db "simple_bank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createTransferRequest struct {
	FromAccountID int64  `json:"from_account" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currencty     string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req createTransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.validAccount(ctx, req.FromAccountID, req.Currencty) {
		return
	}

	if !server.validAccount(ctx, req.ToAccountID, req.Currencty) {
		return
	}

	arg := db.TransferTxParams{FromAccountID: req.FromAccountID, ToAccountID: req.ToAccountID, Amount: req.Amount}
	transfer, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, transfer)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currecny string) bool {
	account, err := server.store.GetAcount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return false
	}

	if account.Currencty != currecny {
		err := fmt.Errorf("account [%d] currency mismatch %s vs %s", account.ID, account.Currencty, currecny)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}
	return true
}
