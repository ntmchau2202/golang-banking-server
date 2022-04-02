package controller

import (
	"bankserver/entity/factory"
	"bankserver/utils"
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
	savingsProduct, err := factory.GetNewSavingsProductFactory().GetSavingsProductByName(savingType)
	if err != nil {
		return "", errors.New("an error occurred when fetching product information")
	}
	cust, err := factory.NewCustomerFactory().GetCustomerByPhone(customerPhone)
	if err != nil {
		return "", errors.New("an error occurred when fetching customer information")
	}

	mtx.Lock()
	savingsAccountID++
	savingsAccountIDStr = strconv.FormatInt(int64(savingsAccountID), 10)
	mtx.Unlock()

	curTime := utils.GetCurrentTimeFormatted()

	// TODO: connect to the blockchain to save transaction
	// If successfull, update true to database
	// else, update pending to database
	// TODO: create a worker to automatically checks for transaction status on blockchain

	// TODO: save to database

	return savingsAccountIDStr, nil
}
