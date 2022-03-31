package controller

type SettleSavingsAccountController struct {
}

func NewSettleSavingsAccountController() *SettleSavingsAccountController {
	return &SettleSavingsAccountController{}
}

func SettleSavingsAccount(
	customerPhone string,
	bankAccount string,
	savingsAccount string,
) (success bool, err error) {
	// TODO: calculate real settle amount
	// TODO: connect to the blockchain to save transaction
	// if successfull, change db status to ok
	// else, change to pending
	// the worker will perform check frequently to see if there is changes/
	// after 3 times having no changes => mark as failed
	return true, nil
}
