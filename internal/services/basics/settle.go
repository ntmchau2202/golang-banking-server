package basic

import (
	"core-banking-server/internal/database"
	"core-banking-server/internal/factory"
	"core-banking-server/internal/models/savingsaccount"
)

type SettleSavingsAccountController struct {
}

func NewSettleSavingsAccountController() *SettleSavingsAccountController {
	return &SettleSavingsAccountController{}
}

func (c *SettleSavingsAccountController) SettleSavingsAccount(
	customerPhone string,
	savingsAccountID string,
	actualInterestAmount float64,
	settleTime string,
) (savingsAccount savingsaccount.SavingsAccount, publicKey string, err error) {
	// TODO: process customer phone here
	// FLOW: save into database first

	savingsAccount, err = c.settleSavingsAccount(savingsAccountID, actualInterestAmount, settleTime)
	if err != nil {
		return
	}
	publicKey, err = c.getCustomerPublicKey(customerPhone)
	return
}

func (c *SettleSavingsAccountController) getCustomerPublicKey(customerPhone string) (key string, err error) {
	cust, err := factory.NewCustomerFactory().GetCustomerByPhone(customerPhone)
	if err != nil {
		return
	}
	return cust.CustomerPublicKey, nil
}

func (c *SettleSavingsAccountController) settleSavingsAccount(
	savingsAccount string,
	actualInterestAmount float64,
	settleTime string,
) (sAcc savingsaccount.SavingsAccount, err error) {
	db, err := database.GetDBConnection()
	if err != nil {
		return
	}

	sAcc, err = db.GetSavingsAccountByID(savingsAccount)
	if err != nil {
		return
	}

	if sAcc.ActualInterestAmount == 0 {
		err = db.SaveSettleSavingsAccount(savingsAccount, settleTime, actualInterestAmount, "")
		if err != nil {
			return
		}

		targetBankAccount, err := db.GetBankAccountByID(sAcc.BankAccountID)
		if err != nil {
			return sAcc, err
		}

		newBalance := targetBankAccount.Balance + actualInterestAmount + sAcc.SavingsAmount

		err = db.UpdateAccountBalance(targetBankAccount.BankAccountID, newBalance)

	}
	return sAcc, err
}
