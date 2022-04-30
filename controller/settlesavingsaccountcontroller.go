package controller

import (
	"bankserver/database"
	"bankserver/entity/signer"
	"bankserver/utils"
	"encoding/json"
	"errors"
)

type SettleSavingsAccountController struct {
}

func NewSettleSavingsAccountController() *SettleSavingsAccountController {
	return &SettleSavingsAccountController{}
}

func (c *SettleSavingsAccountController) SettleSavingsAccount(
	customerPhone string,
	savingsAccountID string,
) (signature string, err error) {
	// TODO: process customer phone here
	// FLOW: save into database first
	currentTime := utils.GetCurrentTimeFormatted()
	err = c.settleSavingsAccount(savingsAccountID, currentTime)
	if err != nil {
		return
	}
	// do something to notify to the client

	// ...
	// request to the blockchain
	db, err := database.GetDBConnection()
	if err != nil {
		return "", errors.New("cannot perform authentication request to blockchain at the moment")
	}
	acc, err := db.GetSavingsAccountByID(savingsAccountID)
	if err != nil {
		return "", errors.New("cannot perform authentication request to blockchain at the moment")
	}
	signer, err := signer.NewSigner("")
	if err != nil {
		return
	}

	data, err := json.Marshal(acc)
	if err != nil {
		return
	}

	signature, err = signer.Sign(string(data))
	return
}

func (c *SettleSavingsAccountController) settleSavingsAccount(
	savingsAccount string,
	settleTime string,
) (err error) {
	db, err := database.GetDBConnection()
	if err != nil {
		return
	}
	// TODO: calculate actual interest amount here
	var actualInterestAmount float64 = 0
	return db.SaveSettleSavingsAccount(savingsAccount, settleTime, actualInterestAmount, "")
}

// func (c *SettleSavingsAccountController) requestSettleConfirmation(
// 	savingsAccountID string,
// 	ownerID string,
// 	ownerPhone string,
// 	settleTime string,
// 	actualInterestAmount float64,
// ) (result response.Response, err error) {
// 	// TODO: hey i forgot the link to the blockchain server :)
// 	var details map[string]interface{} = make(map[string]interface{})
// 	details["savingsaccount_id"] = savingsAccountID
// 	details["owner_id"] = ownerID
// 	details["owner_phone"] = ownerPhone
// 	details["actual_interest_amout"] = strconv.FormatFloat(actualInterestAmount, 'f', -1, 64)
// 	details["settle_time"] = settleTime

// 	msg := request.Request{
// 		Cmd:     message.SETTLE_ONLINE_SAVINGS_ACCOUNT,
// 		Details: details,
// 	}
// 	newClient := client.NewClient("")
// 	result, err = newClient.POST("", msg)
// 	if err != nil {
// 		return
// 	}
// 	return
// }
