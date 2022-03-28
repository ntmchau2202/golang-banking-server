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
	// TODO: modify the database
	return true, nil
}
