package customer

import (
	"bankserver/entity/bankaccount"
	"bankserver/entity/savingsaccount"
)

type Customer struct {
	CustomerID      int64                           `json:"customer_id"`
	CustomerName    string                          `json:"customer_name"`
	CustomerPhone   string                          `json:"customer_phone"`
	BankAccounts    []bankaccount.BankAccount       `json:"bank_accounts"`
	SavingsAccounts []savingsaccount.SavingsAccount `json:"savings_account"`
}
