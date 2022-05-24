package basic

import (
	"core-banking-server/internal/factory"
	"core-banking-server/internal/models/savingsaccount"
	"errors"
)

type getSavingsAccountDetailsController struct {
}

func NewGetSavingsAccountDetailsController() *getSavingsAccountDetailsController {
	return &getSavingsAccountDetailsController{}
}

func (c *getSavingsAccountDetailsController) GetSavingsAccountByID(customerPhone, savingsAccountID string) (acc savingsaccount.SavingsAccount, err error) {
	cust, err := factory.NewCustomerFactory().GetCustomerByPhone(customerPhone)
	if err != nil {
		return acc, errors.New("an error occurred when fetching customer information")
	}

	for _, bankAccount := range cust.BankAccounts {
		for _, savingsAccount := range bankAccount.ListSavingsAccount {
			if savingsAccount.SavingsAccountID == savingsAccountID {
				return savingsAccount, nil
			}
		}
	}

	return savingsaccount.SavingsAccount{}, errors.New("no such savings account with given id associated with customer")
}
