package api

import (
	"bankserver/controller"
	"bankserver/entity/message"
	"bankserver/entity/message/request"
	"bankserver/entity/message/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var msg request.Request
	err := ctx.ShouldBindJSON(msg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("bad request"))
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

	var details map[string]interface{} = make(map[string]interface{})
	details["customer_details"] = cust

	ctx.JSON(http.StatusOK, response.Response{
		Stat:    message.SUCCESS,
		Details: details,
	})
}

func CreateNewSavingsAccount(ctx *gin.Context) {
	var msg request.Request
	err := ctx.ShouldBindJSON(msg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("bad request"))
		return
	}

	customerPhone := msg.Details["customer_phone"].(string)
	bankAccountID := msg.Details["bankAccountID"].(string)
	savingsType := msg.Details["product_type"].(string)
	savingsPeriod := msg.Details["period"].(int)
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

	var details map[string]interface{} = make(map[string]interface{})
	details["new_savings_account"] = newSavingsAccount
	ctx.JSON(http.StatusOK, response.Response{
		Stat:    message.SUCCESS,
		Details: details,
	})
}

func SettleSavingsAccount(ctx *gin.Context) {
	var msg request.Request
	err := ctx.ShouldBindJSON(msg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("bad request"))
		return
	}
	savingsAccountID := msg.Details["savingsaccout_id"].(string)
	customerPhone := msg.Details["customer_phone"].(string)

	if err = validatePhone(customerPhone); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	if err = validateSavingsAccountID(savingsAccountID); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	ctrl := controller.NewSettleSavingsAccountController()
	err = ctrl.SettleSavingsAccount(customerPhone, savingsAccountID)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ErrorResponse(err.Error()))
		return
	}
	var details map[string]interface{} = make(map[string]interface{})
	details["settle_message"] = "success"
	ctx.JSON(http.StatusOK, response.Response{
		Stat:    message.SUCCESS,
		Details: details,
	})
}

func FetchAccountInfo(ctx *gin.Context) {
	var msg request.Request
	err := ctx.ShouldBindJSON(msg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("bad request"))
		return
	}
	customerPhone := msg.Details["customer_phone"].(string)
	bankAccountID := msg.Details["bankAccountID"].(string)

	if err = validatePhone(customerPhone); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	if err = validateBankAccountID(bankAccountID); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	ctrl := controller.NewFetchAccInfController()
	result, err := ctrl.FetchAccInf(customerPhone, bankAccountID)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ErrorResponse(err.Error()))
		return
	}

	var details map[string]interface{} = make(map[string]interface{})

	details["account_info"] = result
	ctx.JSON(http.StatusOK, response.Response{
		Stat:    message.SUCCESS,
		Details: details,
	})
}
