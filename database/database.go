package database

import (
	"bankserver/entity/bankaccount"
	"bankserver/entity/customer"
	"bankserver/entity/savingsaccount"
	"bankserver/entity/savingsproduct"
	"strings"
	"sync"

	"zombiezen.com/go/sqlite"
	db "zombiezen.com/go/sqlite"
)

// singleton
const (
	dbPath string = ""
)

type DatabaseConnection struct {
	dbConn *db.Conn
}

var databaseConnection DatabaseConnection

var singletonDB sync.Once

func GetDBConnection() (conn DatabaseConnection, err error) {
	singletonDB.Do(func() {
		dBConnection, err := db.OpenConn(dbPath, sqlite.OpenReadWrite)
		if err != nil {
			return
		}
		databaseConnection.dbConn = dBConnection
		initConfigs()
	})
	return databaseConnection, nil
}

func initConfigs() (err error) {
	stm := databaseConnection.dbConn.Prep("SELECT * FROM configs")
	if err != nil {
		return
	}
	defer stm.Finalize()
	for {
		if hasRow, err := stm.Step(); err != nil {
			// handle error
		} else if !hasRow {
			break
		}
		field := stm.GetText("field")
		value := stm.GetText("value")
		if strings.Compare(field, "customer_type") == 0 {
			// call customer factory to put customer type here
		} else if strings.Compare(field, "savings_product_type") == 0 {
			// call savings product factory to put savings product type here
		}
	}
	return
}

func (c DatabaseConnection) GetCustomerLoginInfo(customerPhone string) (pwd string, err error) {
	sql := "SELECT * FROM logindetails WHERE customer_phone=$phone"
	stm := c.dbConn.Prep(sql)
	stm.SetText("$phone", customerPhone)
	for {
		if hasRow, err := stm.Step(); err != nil {
			return pwd, err
		} else if !hasRow {
			break
		}
		pwd = stm.GetText("password")
	}
	return
}

func (c DatabaseConnection) GetCustomerByPhone(phone string) (cust customer.Customer, err error) {
	sql := "SELECT * FROM customer WHERE customer_phone=$phone"
	stm := c.dbConn.Prep(sql)
	stm.SetText("$phone", phone)
	for {
		if hasRow, err := stm.Step(); err != nil {
			return cust, err
		} else if !hasRow {
			break
		}
		// create customer instance here
		cust.CustomerType = stm.GetText("customer_type")
		cust.CustomerPhone = stm.GetText("customer_phone")
		cust.CustomerID = stm.GetText("customer_id")
		cust.CustomerName = stm.GetText("customer_name")
	}
	return
}

func (c DatabaseConnection) GetSavingsAccountOfCustomer(customerPhone string) (listSavingsAccount []savingsaccount.SavingsAccount, err error) {
	customer, err := c.GetCustomerByPhone(customerPhone)
	if err != nil {
		return
	}

	sql := "SELECT * FROM customer_savingsaccount WHERE customer_id=$id"
	stm := c.dbConn.Prep(sql)
	stm.SetText("$id", customer.CustomerID)
	for {
		if hasRow, err := stm.Step(); err != nil {
			return listSavingsAccount, err
		} else if !hasRow {
			break
		}
		// fetching..
	}
	return
}

func (c DatabaseConnection) GetBankAccountOfCustomer(customerPhone string) (listBankAccount []bankaccount.BankAccount, err error) {
	customer, err := c.GetCustomerByPhone(customerPhone)
	if err != nil {
		return
	}

	sql := "SELECT * FROM customer_bankaccount WHERE customer_id=$id"
	stm := c.dbConn.Prep(sql)
	stm.SetText("$id", customer.CustomerID)
	for {
		if hasRow, err := stm.Step(); err != nil {
			return listBankAccount, err
		} else if !hasRow {
			break
		}
		listBankAccount = append(listBankAccount, bankaccount.BankAccount{
			OwnerID:       customer.CustomerID,
			BankAccountID: stm.GetText("bankaccount_id"),
			Balance:       stm.GetFloat("bankaccount_balance"),
		})
	}
	return
}

func (c DatabaseConnection) GetSavingsAccountByID(savingsID string) (acc savingsaccount.SavingsAccount, err error) {
	sql := "SELECT * FROM savingsaccout WHERE savingsaccount_id=$id"
	stm := c.dbConn.Prep(sql)
	stm.SetText("$id", savingsID)
	for {
		if hasRow, err := stm.Step(); err != nil {
			return acc, err
		} else if !hasRow {
			break
		}
		acc.SavingsAccountID = stm.GetText("savingsaccount_id")
		acc.SavingsAmount = stm.GetFloat("amount")
		acc.SavingsPeriod = stm.GetInt64("period")
		acc.InterestRate = stm.GetFloat("interest_rate")
		acc.InterestAmount = stm.GetFloat("interest_amount")
		acc.EndTime = stm.GetText("settle_time")
		acc.StartTime = stm.GetText("open_time")
		// TODO: get product type here
		acc.BlockchainConfirmed = stm.GetBool("blockchain_confirmed")
		acc.Currency = stm.GetText("currency")
		acc.SettleInstruction = savingsaccount.SettleType(stm.GetText("settle_instruction"))
	}
	return
}

func (c DatabaseConnection) GetSavingsProductDetails(productName string) (product savingsproduct.SavingsProduct, err error) {
	sql := "SELECT * FROM savingsproduct WHERE product_name=$name"
	stm := c.dbConn.Prep(sql)
	stm.SetText("$name", productName)
	for {
		if hasRow, err := stm.Step(); err != nil {
			return product, err
		} else if !hasRow {
			break
		}
		sqlP := "SELECT * FROM interestrate WHERE product_id=$id"
		stmP := c.dbConn.Prep(sqlP)
		stmP.SetText("$id", stm.GetText("product_id"))
		for {
			if hasRowP, err := stm.Step(); err != nil {
				break
			} else if !hasRowP {
				break
			}
			product.InterestRate[int(stm.GetInt64("period"))] = stm.GetFloat("interest_rate")
		}
		product.ProductName = stm.GetText("product_name")
		product.ProductID = stm.GetText("product_id")
		product.ProductAlias = stm.GetText("product_alias")
	}
	return
}
