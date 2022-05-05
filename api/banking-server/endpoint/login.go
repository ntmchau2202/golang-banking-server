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

	ctrl := basic.NewLoginController()
	cust, err := ctrl.Login(customerPhone, password)
	if err != nil {
		log.Panic("error login:", err)
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.LoginSuccessResponse("login successfully", cust))

}
