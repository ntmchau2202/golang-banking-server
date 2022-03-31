package factory

import (
	"bankserver/database"
	"bankserver/entity/bankaccount"
	"errors"
)

type BankAccountFactory struct {
}

func NewBankAccountFactory() *BankAccountFactory {
	return &BankAccountFactory{}
}

func (c *BankAccountFactory) GetBankAccountsOfCustomer(phone string) (listBankAccount []bankaccount.BankAccount, err error) {
	db, err := database.GetDBConnection()
	if err != nil {
		return listBankAccount, errors.New("unexpected error when getting customer information")
	}

	return db.GetBankAccountOfCustomer(phone)
}

func (c *BankAccountFactory) GetBankAccountByID(bankAccID string) (acc bankaccount.BankAccount, err error) {
	db, err := database.GetDBConnection()
	if err != nil {
		return
	}

	return db.GetBankAccountByID(bankAccID)
}
