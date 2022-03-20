package response

import "bankserver/entity/message"

type Response struct {
	Stat    message.Status         `json:"status"`
	Details map[string]interface{} `json:"details"`
}
