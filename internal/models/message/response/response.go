package response

import (
	"core-banking-server/internal/models/customer"
	"core-banking-server/internal/models/message"
	"core-banking-server/internal/models/savingsaccount"
)

type Response struct {
	Stat    message.Status         `json:"status"`
	Details map[string]interface{} `json:"details"`
}

func ErrorResponse(msg string) (resp Response) {
	resp.Stat = message.ERROR
	resp.Details = make(map[string]interface{})
	resp.Details["message"] = msg
	return
}

func LoginSuccessResponse(msg string, cust customer.Customer) (resp Response) {
	var details map[string]interface{} = make(map[string]interface{})
	details["message"] = msg
	details["customer_details"] = cust
	resp.Stat = message.SUCCESS
	resp.Details = details
	return
}

func FetchAccInfResponse(msg string, cust customer.Customer) (resp Response) {
	var details map[string]interface{} = make(map[string]interface{})
	details["message"] = msg
	details["customer_details"] = cust
	resp.Stat = message.SUCCESS
	resp.Details = details
	return
}

func CreateNewSavingsAccountSuccessResponse(msg string, savingsAccount savingsaccount.SavingsAccount) (resp Response) {
	var details map[string]interface{} = make(map[string]interface{})
	details["message"] = msg
	details["savingsaccount"] = savingsAccount
	resp.Stat = message.SUCCESS
	resp.Details = details
	return
}

func SettleSavingsAccountSuccessResponse(msg string, savingsAccount savingsaccount.SavingsAccount) (resp Response) {
	var details map[string]interface{} = make(map[string]interface{})
	details["message"] = msg
	details["savingsaccount"] = savingsAccount
	resp.Stat = message.SUCCESS
	resp.Details = details
	return
}

func GetSavingsAccountDetailsResponse(msg string, savingsAccount savingsaccount.SavingsAccount) (resp Response) {
	var details map[string]interface{} = make(map[string]interface{})
	details["message"] = msg
	details["savingsaccount"] = savingsAccount
	resp.Stat = message.SUCCESS
	resp.Details = details
	return
}
