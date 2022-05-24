package factory

import (
	"core-banking-server/internal/database"
	"core-banking-server/internal/models/bankaccount"
	"errors"
)

type bankAccountFactory struct {
}

func NewBankAccountFactory() *bankAccountFactory {
	return &bankAccountFactory{}
}

func (c *bankAccountFactory) GetBankAccountsOfCustomer(phone string) (listBankAccount []bankaccount.BankAccount, err error) {
	db, err := database.GetDBConnection()
	if err != nil {
		return listBankAccount, errors.New("unexpected error when getting customer information")
	}

	return db.GetBankAccountOfCustomer(phone)
}

func (c *bankAccountFactory) GetBankAccountByID(bankAccID string) (acc bankaccount.BankAccount, err error) {
	db, err := database.GetDBConnection()
	if err != nil {
		return
	}

	return db.GetBankAccountByID(bankAccID)
}
