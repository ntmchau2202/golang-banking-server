package customer

import (
	"bankserver/entity/bankaccount"
)

type Customer struct {
	CustomerType  string                    `json:"customer_type"`
	CustomerID    string                    `json:"customer_id"`
	CustomerName  string                    `json:"customer_name"`
	CustomerPhone string                    `json:"customer_phone"`
	BankAccounts  []bankaccount.BankAccount `json:"bank_accounts"`
}

var CustomerType []string
