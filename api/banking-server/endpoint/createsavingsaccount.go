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

func CreateSavingsAccount(ctx *gin.Context) {
	var msg request.Request
	err := ctx.ShouldBindJSON(&msg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("bad request"))
		return
	}

	if !msg.CheckCommand(message.CREATE_ONLINE_SAVINGS_ACCOUNT) {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("command mismatch"))
		return
	}

	customerPhone := msg.Details["customer_phone"].(string)
	bankAccountID := msg.Details["bankaccount_id"].(string)
	var savingsAccountID string = ""
	if value, ok := msg.Details["savingsaccount_id"].(string); ok {
		savingsAccountID = value
	}
	savingsType := msg.Details["product_type"].(string)
	savingsPeriod := int(msg.Details["savings_period"].(float64))
	savingsAmount := msg.Details["savings_amount"].(float64)
	interestRate := msg.Details["interest_rate"].(float64)
	estimatedInterestAmount := msg.Details["estimated_interest_amount"].(float64)
	settleInstruction := msg.Details["settle_instruction"].(string)
	currency := msg.Details["currency"].(string)
	openTime := msg.Details["open_time"].(string)

	// checks

	if err = validatePhone(customerPhone); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	if err = validateBankAccountID(bankAccountID); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	if err = validateSavingsType(savingsType); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	if err = validateSavingsPeriod(savingsType, savingsPeriod); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	if err = validateInterestRate(interestRate); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	if err = validateAmount(savingsAmount); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("invalid savings amount"))
		return
	}

	if err = validateAmount(estimatedInterestAmount); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("invalid estimated savings amount"))
		return
	}

	if err = validateSettleInstruction(settleInstruction); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	if err = validateCurrency(currency); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	if err = validateTime(openTime); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	ctrl := basic.NewNewSavingsAccountController()
	account, err := ctrl.CreateNewSavingsAccount(
		customerPhone,
		bankAccountID,
		savingsAccountID,
		savingsType,
		savingsPeriod,
		interestRate,
		savingsAmount,
		estimatedInterestAmount,
		settleInstruction,
		currency,
		openTime,
	)
	if err != nil {
		log.Println("internal server error:", err)
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.CreateNewSavingsAccountSuccessResponse("new savings account created successfully, waiting for blockchain confirmation", account))
}
