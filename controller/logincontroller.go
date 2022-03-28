package controller

import (
	"bankserver/database"
	"errors"
	"strings"
)

type LoginController struct {
}

func NewLoginController() (c *LoginController) {
	return &LoginController{}
}

func (c *LoginController) Authenticate(phone string, password string) (success bool, err error) {
	db, err := database.GetDBConnection()
	if err != nil {
		return false, errors.New("internal server error when performing authentication")
	}

	savedPwd, err := db.GetCustomerLoginInfo(phone)
	if err != nil {
		return false, errors.New("internal server error when performing authentication")
	}
	if strings.Compare(savedPwd, password) == 0 {
		return true, nil
	} else {
		return false, nil
	}
}
