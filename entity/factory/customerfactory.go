package factory

import (
	"bankserver/database"
	"bankserver/entity/customer"
	"strings"
)

type CustomerFactory struct{}

var CustomerType []string

func NewCustomerFactory() *CustomerFactory {
	return &CustomerFactory{}
}

func isCustomerTypeExist(customerType string) bool {
	for _, name := range CustomerType {
		if strings.Compare(customerType, name) == 0 {
			return true
		}
	}
	return false
}

func (f CustomerFactory) PutCustomerType(name string) {
	for _, name := range CustomerType {
		if isCustomerTypeExist(name) {
			return
		}
	}
	CustomerType = append(CustomerType, name)
}

func (f CustomerFactory) GetCustomerByPhone(phone string) (cust customer.Customer, err error) {
	db, err := database.GetDBConnection()
	if err != nil {
		return cust, err
	}

	cust, err = db.GetCustomerByPhone(phone)
	bankAccount, err := GetBankAccountsOfCustomer(cust.CustomerID)
	if err != nil {
		return cust, err
	}

	cust.BankAccounts = append(cust.BankAccounts, bankAccount...)

	// savingsAccount, err := db.GetSavingsAccountOfCustomer(cust.CustomerID)
	// if err != nil {
	// 	return cust, err
	// }

	// cust.SavingsAccounts = append(cust.SavingsAccounts, savingsAccount...)
	return
}
