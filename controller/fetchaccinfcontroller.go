package controller

import (
	"bankserver/entity/bankaccount"
	"bankserver/entity/factory"
)

type FetchAccInfController struct {
}

func NewFetchAccInfController() *FetchAccInfController {
	return &FetchAccInfController{}
}

func (c *FetchAccInfController) FetchAccInf(
	customerPhone string,
	bankAccountID ...string,
) (listBankAcc []bankaccount.BankAccount, err error) {
	if len(bankAccountID) != 0 {
		for _, acc := range bankAccountID {
			bankAcc, err := factory.NewBankAccountFactory().GetBankAccountByID(acc)
			if err != nil {
				continue
			}
			listBankAcc = append(listBankAcc, bankAcc)
		}
	}
	return
}
