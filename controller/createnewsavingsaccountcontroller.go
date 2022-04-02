package controller

import (
	"bankserver/database"
	"bankserver/entity/client"
	"bankserver/entity/factory"
	"bankserver/entity/message"
	"bankserver/entity/message/request"
	"bankserver/entity/message/response"
	"bankserver/entity/savingsaccount"
	"bankserver/utils"
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
	savingType string,
	savingPeriod int,
	savingsAmount float64,
	estimatedInterestAmount float64,
	settleInstruction string,
	currency string,
) (err error) {
	// Flow: create new account and save to database first
	savingsAccount, err := c.createNewAccount(
		customerPhone,
		bankAccountID,
		savingType,
		savingPeriod,
		savingsAmount,
		estimatedInterestAmount,
		settleInstruction,
		currency,
	)
	if err != nil {
		return
	}
	_, err = c.requestCreationConfirmation(
		savingsAccount.SavingsAccountID,
		savingsAccount.OwnerID,
		savingsAccount.OwnerPhone,
		savingsAccount.ProductTypeName,
		savingsAccount.SavingsAmount,
		int(savingsAccount.SavingsPeriod),
		savingsAccount.InterestRate,
		savingsAccount.InterestAmount,
		savingsAccount.StartTime,
		savingsAccount.Currency,
	)
	// the work of updating will be put in another worker
	return
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
) (savingsAcc savingsaccount.SavingsAccount, err error) {
	savingsProduct, err := factory.GetNewSavingsProductFactory().GetSavingsProductByName(savingType)
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

	curTime := utils.GetCurrentTimeFormatted()

	// Flow: save to DB first, then ask the blockchain to confirm
	sAcc := savingsaccount.SavingsAccount{
		SavingsAccountID:  savingsAccountIDStr,
		ProductTypeName:   savingsProduct.ProductName,
		BankAccountID:     bankAccountID,
		SavingsAmount:     savingsAmount,
		InterestAmount:    estimatedInterestAmount,
		StartTime:         curTime,
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

	err = db.SaveCreateNewSavingsAccount(sAcc)
	if err != nil {
		return savingsAcc, errors.New("cannot create new account")
	}
	savingsAcc.SavingsAccountID = savingsAccountIDStr
	return savingsAcc, nil
}

func (c *CreateNewSavingsAccountController) requestCreationConfirmation(
	savingsAccountID string,
	ownerID string,
	ownerPhone string,
	productType string,
	savingsAmount float64,
	savingsPeriod int,
	interstRate float64,
	estimatedInterestAmount float64,
	openTime string,
	currency string,
) (result response.Response, err error) {
	// TODO: hey i forgot the link to the blockchain server :)
	var details map[string]interface{} = make(map[string]interface{})
	details["savingsaccount_id"] = savingsAccountID
	details["owner_id"] = ownerID
	details["owner_phone"] = ownerPhone
	details["product_type"] = productType
	details["savings_amount"] = strconv.FormatFloat(savingsAmount, 'f', -1, 64)
	details["savings_period"] = strconv.FormatInt(int64(savingsPeriod), 10)
	details["interest_rate"] = strconv.FormatFloat(interstRate, 'f', -1, 64)
	details["estimated_interest_amount"] = strconv.FormatFloat(estimatedInterestAmount, 'f', -1, 64)
	details["open_time"] = openTime
	details["currency"] = currency

	msg := request.Request{
		Cmd:     message.CREATE_ONLINE_SAVINGS_ACCOUNT,
		Details: details,
	}
	newClient := client.NewClient("")
	result, err = newClient.POST("", msg)
	if err != nil {
		return
	}
	return
}
