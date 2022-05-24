package basic

import (
	"core-banking-server/internal/database"
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
) (savingsAccount savingsaccount.SavingsAccount, err error) {
	// TODO: process customer phone here
	// FLOW: save into database first

	return c.settleSavingsAccount(savingsAccountID, actualInterestAmount, settleTime)
	// if err != nil {
	// 	return
	// }

	// signer, err := signer.NewSigner("0ae14037ea4665f2c0042a5d15ebf3b6510965c5da80be7c681412b271537b75")
	// if err != nil {
	// 	return
	// }

	// type Txn struct {
	// 	CustomerPhone        string  `json:"customer_phone"`
	// 	SavingsAccountID     string  `json:"savingsaccount_id"`
	// 	ActualInterestAmount float64 `json:"actual_interest_amount"`
	// 	SettleTime           string  `json:"settle_time"`
	// }

	// settleTxn := Txn{
	// 	CustomerPhone:        customerPhone,
	// 	SavingsAccountID:     savingsAccountID,
	// 	ActualInterestAmount: actualInterestAmount,
	// 	SettleTime:           settleTime,
	// }

	// data, err := json.Marshal(settleTxn)
	// if err != nil {
	// 	return
	// }

	// signature, err = signer.Sign(string(data))
	// return
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
