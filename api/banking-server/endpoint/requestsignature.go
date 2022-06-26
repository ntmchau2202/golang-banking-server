package endpoint

import (
	"core-banking-server/internal/models/message"
	"core-banking-server/internal/models/message/request"
	"core-banking-server/internal/models/message/response"
	basic "core-banking-server/internal/services/basics"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSavingsAccountInfo(ctx *gin.Context) {
	// message format
	// {
	// 	"command": "REQUEST_SIGNATURE",
	// 	"details": {
	// 		"customer_phone": "0335909144",
	// 		"savingsaccount_id": "N-0001",
	// 	}
	// }
	// response similar to response of request create account
	var msg request.Request
	err := ctx.ShouldBindJSON(&msg)
	if err != nil {
		fmt.Println("Error binding json:", err)
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	if !msg.CheckCommand(message.REQUEST_SIGNATURE) {
		fmt.Println("command mismatch:", err)
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("command mismatch"))
		return
	}

	customerPhone := msg.Details["customer_phone"].(string)
	savingsAccountID := msg.Details["savingsaccount_id"].(string)

	savingsAccount, err := basic.NewGetSavingsAccountDetailsController().GetSavingsAccountByID(customerPhone, savingsAccountID)
	if err != nil {
		fmt.Println("error getting savings account")
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.GetSavingsAccountDetailsResponse("get savings account details successfully", customerPhone, savingsAccount))
}
