package factory

import (
	"core-banking-server/internal/database"
	"core-banking-server/internal/models/savingsaccount"
)

type savingsAccountFactory struct {
}

func NewSavingsAccountFactory() *savingsAccountFactory {
	return &savingsAccountFactory{}
}

func (f *savingsAccountFactory) GetSavingsAccountByID(savingsAccountID string) (acc savingsaccount.SavingsAccount, err error) {
	db, err := database.GetDBConnection()
	if err != nil {
		return
	}

	acc, err = db.GetSavingsAccountByID(savingsAccountID)
	return
}
