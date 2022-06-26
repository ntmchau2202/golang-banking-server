package basic

import (
	"core-banking-server/internal/database"
	"core-banking-server/internal/models/customer"
	"core-banking-server/internal/models/savingsaccount"
)

type batchQuery struct{}

func GetNewBatchQuery() *batchQuery {
	return &batchQuery{}
}

func (b batchQuery) GetAllCustomers() (listCustomer []customer.Customer, err error) {
	db, err := database.GetDBConnection()
	if err != nil {
		return listCustomer, err
	}
	listCustomer, err = db.GetAllCustomer()
	return
}

func (b batchQuery) GetAllSavingsAccounts() (listSavingsAccount []savingsaccount.SavingsAccount, err error) {
	db, err := database.GetDBConnection()
	if err != nil {
		return listSavingsAccount, err
	}
	listSavingsAccount, err = db.GetAllSavingsAccounts()
	return
}
