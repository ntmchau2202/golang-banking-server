package endpoint

import (
	"core-banking-server/internal/models/message"
	"core-banking-server/internal/models/message/request"
	"core-banking-server/internal/models/message/response"
	basic "core-banking-server/internal/services/basics"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveCustomerPublicKey(ctx *gin.Context) {
	var msg request.Request
	err := ctx.ShouldBindJSON(msg)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse("bad request"))
		return
	}

	if !msg.CheckCommand(message.REGISTER_BLOCKCHAIN_SERVICE) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse("command mismatch"))
		return
	}

	customerID := msg.Details["customer_id"].(string)
	customerPublicKey := msg.Details["public_key"].(string)

	err = basic.GetNewRegisterBlockchainServiceController().SaveCustomerPublicKey(customerID, customerPublicKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusNoContent, "")
}
