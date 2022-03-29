package bankaccount

import (
	"bankserver/entity/savingsaccount"
)

type BankAccount struct {
	OwnerID            string                          `json:"customer_id"`
	BankAccountID      string                          `json:"bankaccount_id"`
	Balance            float64                         `json:"balance"`
	ListSavingsAccount []savingsaccount.SavingsAccount `json:"list_savings_accounts"`
}
