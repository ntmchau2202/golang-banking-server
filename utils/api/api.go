package api

import (
	"bankserver/controller"
	"bankserver/entity/message"
	"bankserver/entity/message/request"
	"bankserver/entity/message/response"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var msg request.Request
	err := ctx.ShouldBindJSON(&msg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	if !msg.CheckCommand(message.LOGIN) {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("command mismatch"))
		return
	}

	customerPhone := msg.Details["customer_phone"].(string)
	password := msg.Details["password"].(string)
	if err = validatePhone(customerPhone); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	if err = validatePassword(password); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	ctrl := controller.NewLoginController()
	cust, err := ctrl.Login(customerPhone, password)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.LoginSuccessResponse("login successfully", cust))
}

func CreateNewSavingsAccount(ctx *gin.Context) {
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

	ctrl := controller.NewNewSavingsAccountController()
	newSavingsAccount, err := ctrl.CreateNewSavingsAccount(
		customerPhone,
		bankAccountID,
		savingsType,
		savingsPeriod,
		savingsAmount,
		estimatedInterestAmount,
		settleInstruction,
		currency,
		openTime,
	)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.CreateNewSavingsAccountSuccessResponse("new savings account created successfully, waiting for blockchain confirmation", newSavingsAccount))
}

func SettleSavingsAccount(ctx *gin.Context) {
	var msg request.Request
	err := ctx.ShouldBindJSON(&msg)
	if err != nil {
		log.Panic(err)
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("bad request"))
		return
	}
	if !msg.CheckCommand(message.SETTLE_ONLINE_SAVINGS_ACCOUNT) {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("command mismatch"))
		return
	}
	savingsAccountID := msg.Details["savingsaccount_id"].(string)
	customerPhone := msg.Details["customer_phone"].(string)

	if err = validatePhone(customerPhone); err != nil {
		log.Panic(err)
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	if err = validateSavingsAccountID(savingsAccountID); err != nil {
		log.Panic(err)
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	ctrl := controller.NewSettleSavingsAccountController()
	err = ctrl.SettleSavingsAccount(customerPhone, savingsAccountID)
	if err != nil {
		log.Panic(err)
		ctx.JSON(http.StatusOK, response.ErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SettleSavingsAccountSuccessResponse("settle savings account successfully, waiting for blockchain confirmation"))
}

func FetchAccountInfo(ctx *gin.Context) {
	var msg request.Request
	err := ctx.ShouldBindJSON(&msg)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("bad request"))
		return
	}
	if !msg.CheckCommand(message.FETCH_LIST_SAVINGS_ACCOUNT) {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("command mismatch"))
		return
	}
	customerPhone := msg.Details["customer_phone"].(string)

	if err = validatePhone(customerPhone); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	ctrl := controller.NewFetchAccInfController()
	result, err := ctrl.FetchAccInf(customerPhone)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.FetchAccInfResponse("get account info successfully", result))
}
