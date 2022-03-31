package database

import (
	"bankserver/entity/bankaccount"
	"bankserver/entity/customer"
	"bankserver/entity/savingsaccount"
	"bankserver/entity/savingsproduct"
	"errors"
)

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

func (c DatabaseConnection) GetBankAccountOfCustomer(customerPhone string) (listBankAccount []bankaccount.BankAccount, err error) {
	customer, err := c.GetCustomerByPhone(customerPhone)
	if err != nil {
		return listBankAccount, errors.New("no such customer with given phone number")
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
		acc := bankaccount.BankAccount{
			OwnerID:       customer.CustomerID,
			BankAccountID: stm.GetText("bankaccount_id"),
			Balance:       stm.GetFloat("bankaccount_balance"),
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

// TODO: recheck table and fields of this function
func (c DatabaseConnection) GetSavingsAccountOfBankAccount(bankAccountID string) (listSavingsAccount []savingsaccount.SavingsAccount, err error) {
	sql := "SELECT * FROM customer_bankaccount WHERE bankaccount_id=$id"
	stm := c.dbConn.Prep(sql)
	stm.SetText("$id", bankAccountID)
	for {
		if hasRow, err := stm.Step(); err != nil {
			return listSavingsAccount, err
		} else if !hasRow {
			break
		}
		savingsAcc := savingsaccount.SavingsAccount{
			SavingsAccountID:    stm.GetText("savingsaccount_id"),
			ProductType:         savingsproduct.SavingsProductType[stm.GetText("product_type")],
			BankAccountID:       bankAccountID,
			SavingsAmount:       stm.GetFloat("amount"),
			InterestAmount:      stm.GetFloat("insterest_amount"),
			StartTime:           stm.GetText("open_time"),
			EndTime:             stm.GetText("settle_time"),
			SavingsPeriod:       stm.GetInt64("savings_period"),
			SettleInstruction:   savingsaccount.SettleType(stm.GetText("settle_instruction")),
			OwnerID:             stm.GetText("ownder_id"),
			InterestRate:        stm.GetFloat("interest_rate"),
			BlockchainConfirmed: stm.GetBool("blockchain_confirmed"),
			Currency:            stm.GetText("currency"),
		}
		listSavingsAccount = append(listSavingsAccount, savingsAcc)
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

func (c DatabaseConnection) GetSavingsAccountConfirmStatus(savingsAccountID string) (isConfirmed bool, err error) {
	savingsAccount, err := c.GetSavingsAccountByID(savingsAccountID)
	if err != nil {
		return false, err
	}

	return savingsAccount.BlockchainConfirmed, nil
}
