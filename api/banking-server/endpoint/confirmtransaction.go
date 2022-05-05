package endpoint

import (
	"core-banking-server/internal/models/message"
	"core-banking-server/internal/models/message/request"
	"core-banking-server/internal/models/message/response"
	basic "core-banking-server/internal/services/basics"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfirmTransaction(ctx *gin.Context) {
	var msg request.Request
	err := ctx.ShouldBindJSON(&msg)
	if err != nil {
		log.Panic("error binding:", err)
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("bad request"))
		return
	}
	if !msg.CheckCommand(message.CONFIRM_TRANSACTION) {
		log.Panic("command mismatch")
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("command mismatch"))
		return
	}

	txnHash := msg.Details["txn_hash"].(string)
	savingsAccountID := msg.Details["savingsaccount_id"].(string)
	action := msg.Details["action"].(string)

	if action != message.CREATE_ONLINE_SAVINGS_ACCOUNT.ToString() && action != message.SETTLE_ONLINE_SAVINGS_ACCOUNT.ToString() {
		log.Panic("invalid action")
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("invalid action"))
		return
	}

	ctrl := basic.NewConfirmTransactionController()
	if action == message.CREATE_ONLINE_SAVINGS_ACCOUNT.ToString() {
		if err := ctrl.SaveOpenTransaction(savingsAccountID, txnHash); err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
			return
		} else {
			ctx.JSON(http.StatusNoContent, struct{}{})
			return
		}
	} else if action == message.SETTLE_ONLINE_SAVINGS_ACCOUNT.ToString() {
		fmt.Println("Going to settle it here")
		if err := ctrl.SaveSettleTransaction(savingsAccountID, txnHash); err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
			return
		} else {
			ctx.JSON(http.StatusNoContent, struct{}{})
			return
		}
	}
}
