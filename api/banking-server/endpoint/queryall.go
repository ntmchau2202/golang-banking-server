package endpoint

import (
	"core-banking-server/internal/models/message/response"
	basic "core-banking-server/internal/services/basics"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryAll(ctx *gin.Context) {
	fmt.Println("Got here")
	params := ctx.Query("topic") // topic=users; topic=savingsaccounts
	fmt.Println("What is our params?:", params)
	if len(params) == 0 {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("missing query params"))
		return
	}
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Content-Type", "application/json")
	fmt.Println("Is everything daijobu?")
	if params == "users" {
		listUsers, err := basic.GetNewBatchQuery().GetAllCustomers()
		if err != nil {
			fmt.Print(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, response.GetAllCustomerResponse("get list users successfully", listUsers))
		return
	} else if params == "savingsaccounts" {
		listSavingsAccount, err := basic.GetNewBatchQuery().GetAllSavingsAccounts()
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, response.GetAllSavingsAccountsResponse("get list savings account successfully", listSavingsAccount))
		return
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse("invalid query"))
		return
	}
}
