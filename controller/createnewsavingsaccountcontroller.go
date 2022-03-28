package controller

import (
	"bankserver/entity/customer"
	"bankserver/entity/savingsproduct"
	"errors"
	"strconv"
	"sync"
)

type CreateNewSavingsAccountController struct {
}

var mtx sync.Mutex
var savingsAccountID = 0

func NewNewSavingsAccountController() (c *CreateNewSavingsAccountController) {
	return &CreateNewSavingsAccountController{}
}

func (c *CreateNewSavingsAccountController) CreateNewAccount(
	customerPhone string,
	savingType string,
	savingPeriod int,
	savingsAmount float64,
	estimatedInterestAmount float64,
	settleInstruction string,
) (savingsAccountIDStr string, err error) {
	savingsProduct, err := savingsproduct.GetNewSavingsProductFactory().GetSavingsProductByName(savingType)
	if err != nil {
		return "", errors.New("an error occurred when fetching product information")
	}
	cust, err := customer.NewCustomerFactory().GetCustomerByPhone(customerPhone)
	if err != nil {
		return "", errors.New("an error occurred when fetching customer information")
	}

	mtx.Lock()
	savingsAccountID++
	savingsAccountIDStr = strconv.FormatInt(int64(savingsAccountID), 10)
	mtx.Unlock()

	// TODO: save to database

	return savingsAccountIDStr, nil
}
