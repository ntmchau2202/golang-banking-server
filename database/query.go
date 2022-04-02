package database

import (
	"bankserver/entity/bankaccount"
	"bankserver/entity/customer"
	"bankserver/entity/savingsaccount"
	"bankserver/entity/savingsproduct"
	"errors"
	"fmt"
)

// OK
func (c DatabaseConnection) GetCustomerLoginInfo(customerPhone string) (pwd string, err error) {
	sql := "SELECT * FROM logindetails WHERE customer_phone=$phone"
	stm := c.dbConn.Prep(sql)
	stm.SetText("$phone", customerPhone)
	if hasRow, err := stm.Step(); err != nil {
		return pwd, err
	} else if !hasRow {
		return pwd, errors.New("no record for given customer phone")
	}
	pwd = stm.GetText("password")
	return
}

// OK
func (c DatabaseConnection) GetCustomerByPhone(phone string) (cust customer.Customer, err error) {
	sql := "SELECT * FROM customer WHERE customer_phone=$phone"
	stm := c.dbConn.Prep(sql)
	stm.SetText("$phone", phone)
	if hasRow, err := stm.Step(); err != nil {
		return cust, err
	} else if !hasRow {
		return cust, errors.New("no customer found")
	}
	// create customer instance here
	cust.CustomerType = stm.GetText("customer_type")
	cust.CustomerPhone = stm.GetText("customer_phone")
	cust.CustomerID = stm.GetText("customer_id")
	cust.CustomerName = stm.GetText("customer_name")

	bankAccounts, err := c.GetBankAccountOfCustomer(cust.CustomerPhone)
	if err != nil {
		return cust, err
	}
	cust.BankAccounts = append(cust.BankAccounts, bankAccounts...)
	return cust, nil
}

// OK
func (c DatabaseConnection) GetBankAccountOfCustomer(customerPhone string) (listBankAccount []bankaccount.BankAccount, err error) {
	sql := "SELECT * FROM customer_bankaccount WHERE customer_phone=$phone"
	stm := c.dbConn.Prep(sql)
	stm.SetText("$phone", customerPhone)
	for {
		if hasRow, err := stm.Step(); err != nil {
			return listBankAccount, err
		} else if !hasRow {
			break
		}
		sqlB := "SELECT * FROM bankaccount WHERE bankaccount_id=$id"
		stmB := c.dbConn.Prep(sqlB)
		stmB.SetText("$id", stm.GetText("bankaccount_id"))
		if ok, err := stmB.Step(); err != nil {
			continue
		} else if !ok {
			continue
		}

		acc := bankaccount.BankAccount{
			OwnerPhone:    customerPhone,
			BankAccountID: stm.GetText("bankaccount_id"),
			Balance:       stmB.GetFloat("bankaccount_balance"),
		}

		// query savings account
		listSavingsAccount, err := c.GetSavingsAccountOfBankAccount(acc.BankAccountID)
		if err != nil {
			return listBankAccount, err
		}
		acc.ListSavingsAccount = append(acc.ListSavingsAccount, listSavingsAccount...)
		listBankAccount = append(listBankAccount, acc)

	}
	return
}

// OK
func (c DatabaseConnection) GetSavingsAccountOfBankAccount(bankAccountID string) (listSavingsAccount []savingsaccount.SavingsAccount, err error) {
	fmt.Println("Start get savings account")
	sql := "SELECT * FROM bankaccount_savingsaccount WHERE bankaccount_id=$id"
	stm := c.dbConn.Prep(sql)
	stm.SetText("$id", bankAccountID)
	for {
		if hasRow, err := stm.Step(); err != nil {
			return listSavingsAccount, err
		} else if !hasRow {
			break
		}
		savingsaccountID := stm.GetText("savingsaccount_id")
		// fmt.Println(savingsaccountID)
		// sqlS := "SELECT * FROM savingsaccount WHERE savingsaccount_id=$id"
		// stmS := c.dbConn.Prep(sqlS)
		// stmS.SetText("$id", savingsaccountID)
		// if ok, err := stmS.Step(); err != nil {
		// 	continue
		// } else if !ok {
		// 	continue
		// }
		// savingsAcc := savingsaccount.SavingsAccount{
		// 	SavingsAccountID:    savingsaccountID,
		// 	ProductType:         savingsproduct.SavingsProductType[stmS.GetText("type")],
		// 	BankAccountID:       bankAccountID,
		// 	SavingsAmount:       stmS.GetFloat("amount"),
		// 	InterestAmount:      stmS.GetFloat("insterest_amount"),
		// 	StartTime:           stmS.GetText("open_time"),
		// 	EndTime:             stmS.GetText("settle_time"),
		// 	SavingsPeriod:       stmS.GetInt64("savings_period"),
		// 	SettleInstruction:   savingsaccount.SettleType(stmS.GetText("settle_instruction")),
		// 	InterestRate:        stmS.GetFloat("interest_rate"),
		// 	BlockchainConfirmed: stmS.GetBool("blockchain_confirmed"),
		// 	Currency:            stmS.GetText("currency"),
		// }
		savingsAcc, err := c.GetSavingsAccountByID(savingsaccountID)
		if err != nil {
			continue
		}
		listSavingsAccount = append(listSavingsAccount, savingsAcc)
	}
	return
}

// OK
func (c DatabaseConnection) GetSavingsAccountByID(savingsID string) (acc savingsaccount.SavingsAccount, err error) {
	sql := "SELECT * FROM savingsaccount WHERE savingsaccount_id=$id"
	stm := c.dbConn.Prep(sql)
	stm.SetText("$id", savingsID)
	if hasRow, err := stm.Step(); err != nil {
		return acc, err
	} else if !hasRow {
		return acc, errors.New("no savings account with id given")
	}
	acc.SavingsAccountID = stm.GetText("savingsaccount_id")
	acc.SavingsAmount = stm.GetFloat("amount")
	acc.SavingsPeriod = stm.GetInt64("period")
	acc.InterestRate = stm.GetFloat("interest_rate")
	acc.InterestAmount = stm.GetFloat("interest_amount")
	acc.EndTime = stm.GetText("settle_time")
	acc.StartTime = stm.GetText("open_time")
	// TODO: get product type here
	acc.CreationConfirmed = stm.GetBool("creation_confirmed")
	acc.SettleConfirmed = stm.GetBool("settle_confirmed")
	acc.Currency = stm.GetText("currency")
	acc.SettleInstruction = savingsaccount.SettleType(stm.GetText("settle_instruction"))
	return
}

func (c DatabaseConnection) GetBankAccountByID(bankAccID string) (bankAcc bankaccount.BankAccount, err error) {
	sql := "SELECT * FROM bankaccount WHERE bankaccount_id=$id"
	stm := c.dbConn.Prep(sql)
	stm.SetText("$id", bankAccID)
	if hasRow, err := stm.Step(); err != nil {
		return bankAcc, err
	} else if !hasRow {
		return bankAcc, errors.New("no bank account with id given")
	}

	bankAcc.Balance = stm.GetFloat("balance")
	bankAcc.BankAccountID = bankAccID
	listSavingsAccount, err := c.GetSavingsAccountOfBankAccount(bankAccID)
	if err != nil {
		return bankAcc, err
	}
	bankAcc.ListSavingsAccount = append(bankAcc.ListSavingsAccount, listSavingsAccount...)
	return
}

// OK
func (c DatabaseConnection) GetSavingsProductDetails(productName string) (product savingsproduct.SavingsProduct, err error) {
	sql := "SELECT * FROM savingsproduct WHERE product_name=$name"
	stm := c.dbConn.Prep(sql)
	stm.SetText("$name", productName)
	product.InterestRate = make(map[int]float64)
	if hasRow, err := stm.Step(); err != nil {
		return product, err
	} else if !hasRow {
		return product, errors.New("no savings product with name given")
	}
	sqlP := "SELECT * FROM interestrate WHERE product_id=$id"
	stmP := c.dbConn.Prep(sqlP)
	stmP.SetText("$id", stm.GetText("product_id"))
	for {
		if hasRowP, err := stmP.Step(); err != nil {
			break
		} else if !hasRowP {
			break
		}
		product.InterestRate[int(stmP.GetInt64("period"))] = stmP.GetFloat("interest_rate")
	}
	product.ProductName = stm.GetText("product_name")
	product.ProductID = stm.GetText("product_id")
	product.ProductAlias = stm.GetText("product_alias")
	return
}

func (c DatabaseConnection) GetSavingsAccountCreationConfirmStatus(savingsAccountID string) (isConfirmed bool, err error) {
	savingsAccount, err := c.GetSavingsAccountByID(savingsAccountID)
	if err != nil {
		return false, err
	}

	return savingsAccount.CreationConfirmed, nil
}

func (c DatabaseConnection) GetSavingsAccountSettleConfirmStatus(savingsAccountID string) (isConfirmed bool, err error) {
	savingsAccount, err := c.GetSavingsAccountByID(savingsAccountID)
	if err != nil {
		return false, err
	}

	return savingsAccount.SettleConfirmed, nil
}
