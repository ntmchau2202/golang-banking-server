package factory

import (
	"core-banking-server/internal/database"
	"core-banking-server/internal/models/customer"
	"strings"
)

type customerFactory struct{}

func NewCustomerFactory() *customerFactory {
	return &customerFactory{}
}

func isCustomerTypeExist(customerType string) bool {
	for _, name := range customer.CustomerType {
		if strings.Compare(customerType, name) == 0 {
			return true
		}
	}
	return false
}

func (f customerFactory) putCustomerType(name string) {
	for _, name := range customer.CustomerType {
		if isCustomerTypeExist(name) {
			return
		}
	}
	customer.CustomerType = append(customer.CustomerType, name)
}

func (f customerFactory) GetCustomerByPhone(phone string) (cust customer.Customer, err error) {
	db, err := database.GetDBConnection()
	if err != nil {
		return cust, err
	}

	cust, err = db.GetCustomerByPhone(phone)
	return
}
