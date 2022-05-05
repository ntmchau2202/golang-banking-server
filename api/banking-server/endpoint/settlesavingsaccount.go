package endpoint

import (
	"core-banking-server/internal/models/message"
	"core-banking-server/internal/models/message/request"
	"core-banking-server/internal/models/message/response"
	basic "core-banking-server/internal/services/basics"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SettleSavingsAccount(ctx *gin.Context) {
	var msg request.Request
	err := ctx.ShouldBindJSON(&msg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("bad request"))
		return
	}
	if !msg.CheckCommand(message.SETTLE_ONLINE_SAVINGS_ACCOUNT) {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("command mismatch"))
		return
	}
	savingsAccountID := msg.Details["savingsaccount_id"].(string)
	customerPhone := msg.Details["customer_phone"].(string)
	settleTime := msg.Details["settle_time"].(string)
	actualInterestAmount := msg.Details["actual_interest_amount"].(float64)

	if err = validatePhone(customerPhone); err != nil {
		log.Panic(err)
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	// if err = validateSavingsAccountID(savingsAccountID); err != nil {
	// 	log.Panic(err)
	// 	ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
	// 	return
	// }

	if err = validateTime(settleTime); err != nil {
		log.Panic(err)
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	ctrl := basic.NewSettleSavingsAccountController()
	account, err := ctrl.SettleSavingsAccount(customerPhone, savingsAccountID, actualInterestAmount, settleTime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SettleSavingsAccountSuccessResponse("settle savings account successfully, waiting for blockchain confirmation", account))
}
