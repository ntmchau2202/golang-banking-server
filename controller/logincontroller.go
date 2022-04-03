package controller

import (
	"bankserver/database"
	"bankserver/entity/customer"
	"bankserver/entity/factory"
	"errors"
	"strings"
)

type LoginController struct {
}

func NewLoginController() (c *LoginController) {
	return &LoginController{}
}

func (c *LoginController) authenticate(phone string, password string) (success bool, err error) {
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

func (c *LoginController) Login(phone string, password string) (cust customer.Customer, err error) {
	if ok, err := c.authenticate(phone, password); err != nil {
		return cust, err
	} else if !ok {
		return cust, errors.New("invalid password")
	}

	cust, err = factory.NewCustomerFactory().GetCustomerByPhone(phone)
	if err != nil {
		return cust, errors.New("cannot get customer informaton")
	}
	return
}
