package bankaccount

import (
	"bankserver/entity/customer"
	"bankserver/entity/savingsaccount"
)

type BankAccount struct {
	Owner               customer.Customer               `json:"customer"`
	BankAccountID       string                          `json:"bankaccount_id"`
	Balance             float64                         `json:"balance"`
	ListSavingsAccounts []savingsaccount.SavingsAccount `json:"list_savings_accounts`
}
