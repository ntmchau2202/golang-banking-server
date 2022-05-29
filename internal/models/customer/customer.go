package customer

import "core-banking-server/internal/models/bankaccount"

type Customer struct {
	CustomerType      string                    `json:"customer_type"`
	CustomerID        string                    `json:"customer_id"`
	CustomerName      string                    `json:"customer_name"`
	CustomerPhone     string                    `json:"customer_phone"`
	CustomerPublicKey string                    `json:"public_key"`
	BankAccounts      []bankaccount.BankAccount `json:"bank_accounts"`
}

var CustomerType []string
