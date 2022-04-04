package controller

import (
	"bankserver/entity/customer"
	"bankserver/entity/factory"
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
