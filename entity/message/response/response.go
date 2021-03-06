package response

import (
	"bankserver/entity/customer"
	"bankserver/entity/message"
)

type Response struct {
	Stat    message.Status         `json:"status"`
	Details map[string]interface{} `json:"details"`
}

func ErrorResponse(msg string) (resp Response) {
	resp.Stat = message.SUCCESS
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

func CreateNewSavingsAccountSuccessResponse(msg string, signature string, id string) (resp Response) {
	var details map[string]interface{} = make(map[string]interface{})
	details["message"] = msg
	details["signature"] = signature
	details["savingsaccount_id"] = id
	resp.Stat = message.SUCCESS
	resp.Details = details
	return
}

func SettleSavingsAccountSuccessResponse(msg string, signature string) (resp Response) {
	var details map[string]interface{} = make(map[string]interface{})
	details["message"] = msg
	details["signature"] = signature
	resp.Stat = message.SUCCESS
	resp.Details = details
	return
}
