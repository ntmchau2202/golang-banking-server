package bankaccount

import "bankserver/entity/customer"

type BankAccount struct {
	Owner         customer.Customer `json:"customer"`
	BankAccountID string            `json:"bankaccount_id"`
	Balance       float64           `json:"balance"`
}
