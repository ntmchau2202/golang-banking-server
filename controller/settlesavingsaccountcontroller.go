package controller

import (
	"bankserver/database"
	"bankserver/entity/signer"
	"encoding/json"
)

type SettleSavingsAccountController struct {
}

func NewSettleSavingsAccountController() *SettleSavingsAccountController {
	return &SettleSavingsAccountController{}
}

func (c *SettleSavingsAccountController) SettleSavingsAccount(
	customerPhone string,
	savingsAccountID string,
	actualInterestAmount float64,
	settleTime string,
) (signature string, err error) {
	// TODO: process customer phone here
	// FLOW: save into database first
	err = c.settleSavingsAccount(savingsAccountID, actualInterestAmount, settleTime)
	if err != nil {
		return
	}
	signer, err := signer.NewSigner("0ae14037ea4665f2c0042a5d15ebf3b6510965c5da80be7c681412b271537b75")
	if err != nil {
		return
	}

	type Txn struct {
		CustomerPhone        string  `json:"customer_phone"`
		SavingsAccountID     string  `json:"savingsaccount_id"`
		ActualInterestAmount float64 `json:"actual_interest_amount"`
		SettleTime           string  `json:"settle_time"`
	}

	settleTxn := Txn{
		CustomerPhone:        customerPhone,
		SavingsAccountID:     savingsAccountID,
		ActualInterestAmount: actualInterestAmount,
		SettleTime:           settleTime,
	}

	data, err := json.Marshal(settleTxn)
	if err != nil {
		return
	}

	signature, err = signer.Sign(string(data))
	return
}

func (c *SettleSavingsAccountController) settleSavingsAccount(
	savingsAccount string,
	actualInterestAmount float64,
	settleTime string,
) (err error) {
	db, err := database.GetDBConnection()
	if err != nil {
		return
	}

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
