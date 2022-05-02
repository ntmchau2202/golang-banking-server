package controller

import (
	"bankserver/database"
	"bankserver/entity/factory"
	"bankserver/entity/savingsaccount"
	"bankserver/entity/signer"
	"encoding/json"
	"errors"
	"fmt"
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
	savingType string,
	savingPeriod int,
	savingsAmount float64,
	estimatedInterestAmount float64,
	settleInstruction string,
	currency string,
	openTime string,
) (savingsAccountID string, signature string, err error) {
	// Flow: create new account and save to database first
	// TODO: process open time here
	savingsAccount, err := c.createNewAccount(
		customerPhone,
		bankAccountID,
		savingType,
		savingPeriod,
		savingsAmount,
		estimatedInterestAmount,
		settleInstruction,
		currency,
		openTime,
	)
	if err != nil {
		return
	}
	// create a json message and sign it with the bank's private key
	// then return this to the customer
	signer, err := signer.NewSigner("0ae14037ea4665f2c0042a5d15ebf3b6510965c5da80be7c681412b271537b75")
	if err != nil {
		return
	}

	data, err := json.Marshal(savingsAccount)
	if err != nil {
		return
	}

	fmt.Println(string(data))

	signature, err = signer.Sign(string(data))
	return savingsAccount.SavingsAccountID, signature, err
}

func (c *CreateNewSavingsAccountController) createNewAccount(
	customerPhone string,
	bankAccountID string,
	savingType string,
	savingPeriod int,
	savingsAmount float64,
	estimatedInterestAmount float64,
	settleInstruction string,
	currency string,
	openTime string,
) (savingsAcc savingsaccount.SavingsAccount, err error) {
	savingsProduct, err := factory.NewSavingsProductFactory().GetSavingsProductByName(savingType)
	if err != nil {
		return savingsAcc, errors.New("an error occurred when fetching product information")
	}
	cust, err := factory.NewCustomerFactory().GetCustomerByPhone(customerPhone)
	if err != nil {
		return savingsAcc, errors.New("an error occurred when fetching customer information")
	}

	mtx.Lock()
	savingsAccountID++
	// TODO: create proper savings account id here
	savingsAccountIDStr := strconv.FormatInt(int64(savingsAccountID), 10)
	mtx.Unlock()

	// curTime := utils.GetCurrentTimeFormatted()

	// Flow: save to DB first, then ask the blockchain to confirm
	savingsAcc = savingsaccount.SavingsAccount{
		SavingsAccountID:  savingsAccountIDStr,
		ProductTypeName:   savingsProduct.ProductName,
		BankAccountID:     bankAccountID,
		SavingsAmount:     savingsAmount,
		InterestAmount:    estimatedInterestAmount,
		StartTime:         openTime,
		SavingsPeriod:     int64(savingPeriod),
		SettleInstruction: savingsaccount.SettleType(settleInstruction),
		OwnerID:           cust.CustomerID,
		InterestRate:      savingsProduct.InterestRate[savingPeriod],
		Currency:          currency,
	}
	db, err := database.GetDBConnection()
	if err != nil {
		return savingsAcc, errors.New("cannot connect to database")
	}

	err = db.SaveCreateNewSavingsAccount(savingsAcc)
	if err != nil {
		return savingsAcc, errors.New("cannot create new account")
	}
	savingsAcc.SavingsAccountID = savingsAccountIDStr
	return savingsAcc, nil
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
