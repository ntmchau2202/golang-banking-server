package controller

import "bankserver/entity/savingsaccount"

type GetSavingsAccountOnChain struct {
}

func NewGetSavingsAccountOnChain() *GetSavingsAccountOnChain {
	return &GetSavingsAccountOnChain{}
}

func (c *GetSavingsAccountOnChain) GetSavingsAccountOnChain(savingsAccountID string) (acc savingsaccount.SavingsAccount, err error) {
	// TODO: should we make this into our desired worker?
	return
}
