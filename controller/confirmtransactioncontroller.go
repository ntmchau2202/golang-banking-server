package controller

import (
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

	// check savings account id
	return
}

func (c *ConfirmTransactionController) SaveSettleTransaction(savingsAccountID string, txnHash string) (err error) {
	if !strings.HasPrefix(txnHash, "0x") {
		return errors.New("invalid transaction hash")
	}

	// check savings account id
	return
}
