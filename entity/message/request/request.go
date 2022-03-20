package request

import "bankserver/entity/message"

type Request struct {
	Cmd     message.Command        `json:"command"`
	Details map[string]interface{} `json:"details"`
}

func (r Request) CheckCommand(cmd message.Command) (match bool) {
	return r.Cmd.Equals(cmd)
}
