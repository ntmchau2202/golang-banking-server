package basic

import (
	"core-banking-server/internal/factory"
	"core-banking-server/internal/models/customer"
)

type FetchAccInfController struct {
}

func NewFetchAccInfController() *FetchAccInfController {
	return &FetchAccInfController{}
}

func (c *FetchAccInfController) FetchAccInf(
	customerPhone string,
) (cust customer.Customer, err error) {
	// borrow controller :)
	cust, err = factory.NewCustomerFactory().GetCustomerByPhone(customerPhone)
	return
}
