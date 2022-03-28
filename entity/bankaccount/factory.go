package bankaccount

import (
	"bankserver/database"
	"errors"
)

type BankAccountFactory struct {
}

func NewBankAccountFactory() *BankAccountFactory {
	return &BankAccountFactory{}
}

func GetBankAccountsOfCustomer(phone string) (listBankAccount []BankAccount, err error) {
	db, err := database.GetDBConnection()
	if err != nil {
		return listBankAccount, errors.New("unexpected error when getting customer information")
	}

	return db.GetBankAccountOfCustomer(phone)
}

func GetBankAccountByID(bankAccID string) (acc BankAccount, err error) {
	return
}
