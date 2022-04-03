package bankaccount

import (
	"bankserver/entity/savingsaccount"
)

type BankAccount struct {
	OwnerPhone         string                          `json:"customer_phone"`
	BankAccountID      string                          `json:"bankaccount_id"`
	Balance            float64                         `json:"balance"`
	ListSavingsAccount []savingsaccount.SavingsAccount `json:"savings_accounts"`
}
