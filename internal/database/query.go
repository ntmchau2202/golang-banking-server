package database

import (
	"core-banking-server/internal/models/bankaccount"
	"core-banking-server/internal/models/customer"
	"core-banking-server/internal/models/savingsaccount"
	"core-banking-server/internal/models/savingsproduct"
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

func (c DatabaseConnection) GetcustomerByID(id string) (cust customer.Customer, err error) {
	sql := "SELECT * FROM customer WHERE customer_id=$id"
	stm := c.dbConn.Prep(sql)
	stm.SetText("$id", id)
	if hasRow, err := stm.Step(); err != nil {
		return cust, err
	} else if !hasRow {
		return cust, errors.New("no customer found")
	}

	phone := stm.GetText("customer_phone")
	return c.GetCustomerByPhone(phone)
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
	cust.CustomerPublicKey = stm.GetText("public_key")

	bankAccounts, err := c.GetBankAccountOfCustomer(cust.CustomerPhone)
	if err != nil {
		return cust, err
	}
	if len(bankAccounts) == 0 {
		cust.BankAccounts = []bankaccount.BankAccount{}
	} else {
		cust.BankAccounts = append(cust.BankAccounts, bankAccounts...)
	}
	return cust, nil
}

// OK
func (c DatabaseConnection) GetBankAccountOfCustomer(customerPhone string) (listBankAccount []bankaccount.BankAccount, err error) {
	sql := "SELECT DISTINCT * FROM customer_bankaccount WHERE customer_phone=$phone"
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
		var acc bankaccount.BankAccount
		for {
			if ok, err := stmB.Step(); err != nil {
				break
			} else if !ok {
				break
			}
			acc = bankaccount.BankAccount{
				OwnerPhone:    customerPhone,
				BankAccountID: stm.GetText("bankaccount_id"),
				Balance:       stmB.GetFloat("bankaccount_balance"),
			}

			// query savings account
			listSavingsAccount, err := c.GetSavingsAccountOfBankAccount(acc.BankAccountID)
			if err != nil || len(listSavingsAccount) == 0 {
				acc.ListSavingsAccount = []savingsaccount.SavingsAccount{}
			}
			acc.ListSavingsAccount = append(acc.ListSavingsAccount, listSavingsAccount...)
			listBankAccount = append(listBankAccount, acc)
		}
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
		fmt.Println("Get one savingaccountID:", savingsaccountID)
		savingsAcc, err := c.GetSavingsAccountByID(savingsaccountID)
		savingsAcc.BankAccountID = bankAccountID
		fmt.Println(savingsAcc.BankAccountID)
		if err != nil {
			continue
		}
		listSavingsAccount = append(listSavingsAccount, savingsAcc)
	}
	return
}

// OK
func (c DatabaseConnection) GetSavingsAccountByID(savingsID string) (acc savingsaccount.SavingsAccount, err error) {
	sql := "SELECT bankaccount_id FROM bankaccount_savingsaccount WHERE savingsaccount_id=$id"
	stm := c.dbConn.Prep(sql)
	stm.SetText("$id", savingsID)

	if hasRow, err := stm.Step(); err != nil {
		return acc, err
	} else if !hasRow {
		return acc, errors.New("savings account does not associate with any bank account")
	}

	acc.BankAccountID = stm.GetText("bankaccount_id")

	sql = "SELECT * FROM savingsaccount WHERE savingsaccount_id=$id"
	stm = c.dbConn.Prep(sql)
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
	acc.ActualInterestAmount = stm.GetFloat("actual_interest_amount")
	acc.EndTime = stm.GetText("settle_time")
	acc.StartTime = stm.GetText("open_time")
	acc.ProductTypeName = stm.GetText("type")
	acc.CreationConfirmed = stm.GetText("creation_confirmed")
	acc.SettleConfirmed = stm.GetText("settle_confirmed")
	acc.Currency = stm.GetText("currency")
	acc.SettleInstruction = savingsaccount.SettleType(stm.GetText("settle_instruction"))
	acc.ConfirmStatus = stm.GetInt64("status")

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

func (c DatabaseConnection) GetSavingsAccountCreationConfirmStatus(savingsAccountID string) (isConfirmed string, err error) {
	savingsAccount, err := c.GetSavingsAccountByID(savingsAccountID)
	if err != nil {
		return "", err
	}

	return savingsAccount.CreationConfirmed, nil
}

func (c DatabaseConnection) GetSavingsAccountSettleConfirmStatus(savingsAccountID string) (isConfirmed string, err error) {
	savingsAccount, err := c.GetSavingsAccountByID(savingsAccountID)
	if err != nil {
		return "", err
	}

	return savingsAccount.SettleConfirmed, nil
}

func (c DatabaseConnection) IsAccountCreationConfirmed(savingsAccountID string) (isConfirmed bool, txnHash string, err error) {
	sql := "SELECT creation_confirmed FROM savingsaccount WHERE savingsaccount_id=$id"
	stm := c.dbConn.Prep(sql)
	stm.SetText("$id", savingsAccountID)

	if hasRow, err := stm.Step(); err != nil {
		return false, "", err
	} else if !hasRow {
		return false, "", errors.New("no records for this savings account found")
	} else if len(stm.GetText("creation_confirmed")) == 0 {
		return false, "", nil
	}
	return true, stm.GetText("creation_confirmed"), nil
}

func (c DatabaseConnection) IsAccountSettlementConfirmed(savingsAccountID string) (isConfirmed bool, txnHash string, err error) {
	sql := "SELECT settle_confirmed FROM savingsaccount WHERE savingsaccount_id=$id"
	stm := c.dbConn.Prep(sql)
	stm.SetText("$id", savingsAccountID)

	if hasRow, err := stm.Step(); err != nil {
		return false, "", err
	} else if !hasRow {
		return false, "", errors.New("no records for this savings account found")
	} else if len(stm.GetText("settle_confirmed")) == 0 {
		return false, "", nil
	}
	return true, stm.GetText("settle_confirmed"), nil
}

func (c DatabaseConnection) GetCustomerPublicKey(customerID string) (publicKey string, err error) {
	cust, err := c.GetcustomerByID(customerID)
	if err != nil {
		return
	}
	return cust.CustomerPublicKey, nil
}
