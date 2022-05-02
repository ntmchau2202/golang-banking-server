package controller

import (
	"bankserver/database"
	"bankserver/entity/factory"
	"errors"
	"strings"
)

type ConfirmTransactionController struct {
}

func NewConfirmTransactionController() *ConfirmTransactionController {
	return &ConfirmTransactionController{}
}

func (c *ConfirmTransactionController) SaveOpenTransaction(savingsAccountID string, txnHash string) (err error) {
	if !strings.HasPrefix(txnHash, "0x") {
		return errors.New("invalid transaction hash")
	}
	return c.saveOpenTransaction(savingsAccountID, txnHash)
}

func (c *ConfirmTransactionController) saveOpenTransaction(savingsAccountID string, txnHash string) (err error) {
	_, err = factory.NewSavingsAccountFactory().GetSavingsAccountByID(savingsAccountID)
	if err != nil {
		return err
	}

	// save
	db, err := database.GetDBConnection()
	if err != nil {
		return
	}

	return db.SaveSavingAccountCreationConfirmationStatus(savingsAccountID, txnHash)
}

func (c *ConfirmTransactionController) SaveSettleTransaction(savingsAccountID string, txnHash string) (err error) {
	if !strings.HasPrefix(txnHash, "0x") {
		return errors.New("invalid transaction hash")
	}

	// save
	db, err := database.GetDBConnection()
	if err != nil {
		return
	}

	return db.SaveSavingAccountSettleConfirmationStatus(savingsAccountID, txnHash)
}

func (c *ConfirmTransactionController) saveSettleTransaction(savingsAccountID string, txnHash string) (err error) {
	_, err = factory.NewSavingsAccountFactory().GetSavingsAccountByID(savingsAccountID)
	if err != nil {
		return err
	}

	// save
	db, err := database.GetDBConnection()
	if err != nil {
		return
	}

	return db.SaveSavingAccountSettleConfirmationStatus(savingsAccountID, txnHash)
}
