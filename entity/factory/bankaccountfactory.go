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

func GetBankAccountsOfCustomer(phone string) (listBankAccount []bankaccount.BankAccount, err error) {
	db, err := database.GetDBConnection()
	if err != nil {
		return listBankAccount, errors.New("unexpected error when getting customer information")
	}

	return db.GetBankAccountOfCustomer(phone)
}

func GetBankAccountByID(bankAccID string) (acc bankaccount.BankAccount, err error) {
	return
}
