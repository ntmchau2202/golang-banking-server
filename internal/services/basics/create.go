package basic

import (
	"core-banking-server/internal/database"
	"core-banking-server/internal/factory"
	"core-banking-server/internal/models/savingsaccount"
	"errors"
	"strconv"
	"sync"
)

type CreateNewSavingsAccountController struct {
}

var mtx sync.Mutex
var savingsAccountID = 0

func NewNewSavingsAccountController() (c *CreateNewSavingsAccountController) {
	return &CreateNewSavingsAccountController{}
}

func (c *CreateNewSavingsAccountController) CreateNewSavingsAccount(
	customerPhone string,
	bankAccountID string,
	savingsAccountID string,
	savingType string,
	savingPeriod int,
	interestRate float64,
	savingsAmount float64,
	estimatedInterestAmount float64,
	settleInstruction string,
	currency string,
	openTime string,
	// ) (savingsAccountID string, signature string, err error) {
) (savingsAccount savingsaccount.SavingsAccount, customerPublicKey string, err error) {
	// Flow: create new account and save to database first
	// TODO: process open time here
	savingsAccount, err = c.createNewAccount(
		customerPhone,
		bankAccountID,
		savingsAccountID,
		savingType,
		savingPeriod,
		interestRate,
		savingsAmount,
		estimatedInterestAmount,
		settleInstruction,
		currency,
		openTime,
	)
	if err != nil {
		return
	}
	customerPublicKey, err = c.getCustomerPublicKey(customerPhone)
	return
}

func (c *CreateNewSavingsAccountController) getCustomerPublicKey(customerPhone string) (key string, err error) {
	cust, err := factory.NewCustomerFactory().GetCustomerByPhone(customerPhone)
	if err != nil {
		return
	}
	return cust.CustomerPublicKey, nil
}

func (c *CreateNewSavingsAccountController) createNewAccount(
	customerPhone string,
	bankAccountID string,
	currentSavingsAccountID string,
	savingType string,
	savingPeriod int,
	interestRate float64,
	savingsAmount float64,
	estimatedInterestAmount float64,
	settleInstruction string,
	currency string,
	openTime string,
) (savingsAcc savingsaccount.SavingsAccount, err error) {
	// savingsProduct, err := factory.NewSavingsProductFactory().GetSavingsProductByName(savingType)
	// if err != nil {
	// 	return savingsAcc, errors.New("an error occurred when fetching product information")
	// }
	cust, err := factory.NewCustomerFactory().GetCustomerByPhone(customerPhone)
	if err != nil {
		return savingsAcc, errors.New("an error occurred when fetching customer information")
	}

	savingsAcc = savingsaccount.SavingsAccount{
		ProductTypeName:   savingType,
		BankAccountID:     bankAccountID,
		SavingsAmount:     savingsAmount,
		InterestAmount:    estimatedInterestAmount,
		StartTime:         openTime,
		SavingsPeriod:     int64(savingPeriod),
		SettleInstruction: savingsaccount.SettleType(settleInstruction),
		OwnerID:           cust.CustomerID,
		InterestRate:      interestRate,
		Currency:          currency,
	}
	var needUpdate bool = false
	if len(currentSavingsAccountID) == 0 {
		mtx.Lock()
		savingsAccountID++
		// TODO: create proper savings account id here
		savingsAccountIDStr := strconv.FormatInt(int64(savingsAccountID), 10)
		mtx.Unlock()
		currentSavingsAccountID = savingsAccountIDStr
		needUpdate = true
	}

	savingsAcc.SavingsAccountID = currentSavingsAccountID
	if needUpdate {
		db, err := database.GetDBConnection()
		if err != nil {
			return savingsAcc, errors.New("cannot connect to database")
		}

		err = db.SaveCreateNewSavingsAccount(savingsAcc)
		if err != nil {
			return savingsAcc, errors.New("cannot create new account")
		}
		// update account balance here
		// calculate new balance
		var currentBankAccountBalance float64 = 0
		for _, acc := range cust.BankAccounts {
			if acc.BankAccountID == bankAccountID {
				currentBankAccountBalance = acc.Balance
				break
			}
		}

		newBankAccountBalance := currentBankAccountBalance - savingsAmount
		err = db.UpdateAccountBalance(bankAccountID, newBankAccountBalance)
		targetBankAccount, err := db.GetBankAccountByID(savingsAcc.BankAccountID)
		if err != nil {
			return savingsAcc, err
		}
		err = db.AddSavingsAccountToBankAccount(savingsAcc.SavingsAccountID, targetBankAccount.BankAccountID)
	}

	return savingsAcc, err
}

// func (c *CreateNewSavingsAccountController) requestCreationConfirmation(
// 	savingsAccountID string,
// 	ownerID string,
// 	ownerPhone string,
// 	productType string,
// 	savingsAmount float64,
// 	savingsPeriod int,
// 	interstRate float64,
// 	estimatedInterestAmount float64,
// 	openTime string,
// 	currency string,
// ) (result response.Response, err error) {
// 	// TODO: hey i forgot the link to the blockchain server :)
// 	var details map[string]interface{} = make(map[string]interface{})
// 	details["savingsaccount_id"] = savingsAccountID
// 	details["owner_id"] = ownerID
// 	details["owner_phone"] = ownerPhone
// 	details["product_type"] = productType
// 	details["savings_amount"] = strconv.FormatFloat(savingsAmount, 'f', -1, 64)
// 	details["savings_period"] = strconv.FormatInt(int64(savingsPeriod), 10)
// 	details["interest_rate"] = strconv.FormatFloat(interstRate, 'f', -1, 64)
// 	details["estimated_interest_amount"] = strconv.FormatFloat(estimatedInterestAmount, 'f', -1, 64)
// 	details["open_time"] = openTime
// 	details["currency"] = currency

// 	msg := request.Request{
// 		Cmd:     message.CREATE_ONLINE_SAVINGS_ACCOUNT,
// 		Details: details,
// 	}
// 	newClient := client.NewClient("")
// 	result, err = newClient.POST("", msg)
// 	if err != nil {
// 		return
// 	}
// 	return
// }
