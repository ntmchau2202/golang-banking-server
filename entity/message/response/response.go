package response

import "bankserver/entity/message"

type Response struct {
	Stat    message.Status         `json:"status"`
	Details map[string]interface{} `json:"details"`
}

func ErrorResponse(msg string) (resp Response) {
	resp.Stat = message.SUCCESS
	resp.Details["message"] = msg
	return
}
