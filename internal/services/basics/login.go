package basic

import (
	"core-banking-server/internal/database"
	"core-banking-server/internal/factory"
	"core-banking-server/internal/models/customer"
	"errors"
	"fmt"
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
		fmt.Println("error getting database connection:", err)
		return false, errors.New("internal server error when performing authentication")
	}

	savedPwd, err := db.GetCustomerLoginInfo(phone)
	if err != nil {
		if err.Error() == "no record for given customer phone" {
			return true, errors.New("invalid customer")
		}
		return false, errors.New("internal server error when performing authentication")
	}
	if strings.Compare(savedPwd, password) == 0 {
		return true, nil
	} else {
		fmt.Println("invalid password")
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
