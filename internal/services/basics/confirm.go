package basic

import (
	"core-banking-server/internal/database"
	"core-banking-server/internal/factory"
	"errors"
	"strings"
)

type ConfirmTransactionController struct {
}

func NewConfirmTransactionController() *ConfirmTransactionController {
	return &ConfirmTransactionController{}
}

func (c *ConfirmTransactionController) SaveOpenTransaction(savingsAccountID string, txnHash string, ipfsHash string) (err error) {
	if !strings.HasPrefix(txnHash, "0x") {
		return errors.New("invalid transaction hash")
	}
	return c.saveOpenTransaction(savingsAccountID, txnHash, ipfsHash)
}

func (c *ConfirmTransactionController) saveOpenTransaction(savingsAccountID string, txnHash string, ipfsHash string) (err error) {
	_, err = factory.NewSavingsAccountFactory().GetSavingsAccountByID(savingsAccountID)
	if err != nil {
		return err
	}

	// save
	db, err := database.GetDBConnection()
	if err != nil {
		return
	}

	if confirmed, savedTxnHash, err := db.IsAccountCreationConfirmed(savingsAccountID); err != nil {
		return err
	} else if !confirmed {
		return db.SaveSavingAccountCreationConfirmationStatus(savingsAccountID, txnHash, ipfsHash)
	} else {
		return errors.New("transaction has already been confirmed with txn hash " + savedTxnHash)
	}

}

func (c *ConfirmTransactionController) SaveSettleTransaction(savingsAccountID string, txnHash string, ipfsHash string) (err error) {
	if !strings.HasPrefix(txnHash, "0x") {
		return errors.New("invalid transaction hash")
	}

	// save
	db, err := database.GetDBConnection()
	if err != nil {
		return
	}

	if confirmed, savedTxnHash, err := db.IsAccountSettlementConfirmed(savingsAccountID); err != nil {
		return err
	} else if !confirmed {
		return db.SaveSavingAccountSettleConfirmationStatus(savingsAccountID, txnHash, ipfsHash)
	} else {
		return errors.New("transaction has already been confirmed with txn hash " + savedTxnHash)
	}
}

func (c *ConfirmTransactionController) saveSettleTransaction(savingsAccountID string, txnHash string, ipfsHash string) (err error) {
	_, err = factory.NewSavingsAccountFactory().GetSavingsAccountByID(savingsAccountID)
	if err != nil {
		return err
	}

	// save
	db, err := database.GetDBConnection()
	if err != nil {
		return
	}

	return db.SaveSavingAccountSettleConfirmationStatus(savingsAccountID, txnHash, ipfsHash)
}
