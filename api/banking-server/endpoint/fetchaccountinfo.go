package endpoint

import (
	"core-banking-server/internal/models/message"
	"core-banking-server/internal/models/message/request"
	"core-banking-server/internal/models/message/response"
	basic "core-banking-server/internal/services/basics"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchAccountInfo(ctx *gin.Context) {
	var msg request.Request
	err := ctx.ShouldBindJSON(&msg)
	if err != nil {
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

	ctrl := basic.NewFetchAccInfController()
	result, err := ctrl.FetchAccInf(customerPhone)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.FetchAccInfResponse("get account info successfully", result))
}
