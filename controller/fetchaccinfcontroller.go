package controller

import (
	"bankserver/entity/bankaccount"
)

type FetchAccInfController struct {
}

func NewFetchAccInfController() *FetchAccInfController {
	return &FetchAccInfController{}
}

func FetchAccInf(
	customerPhone string,
	bankAccountID ...string,
) (listBankAcc []bankaccount.BankAccount, err error) {
	if len(bankAccountID) != 0 {
		// for _, acc := range bankAccountID {

		// }
	}
	return
}
