package message

import "strings"

type Command string

const (
	LOGIN                         Command = "LOGIN"
	SETTLE_ONLINE_SAVINGS_ACCOUNT Command = "SETTLE_ONLINE_SAVINGS_ACCOUNT"
	CREATE_ONLINE_SAVINGS_ACCOUNT Command = "CREATE_ONLINE_SAVINGS_ACCOUNT"
	FETCH_LIST_SAVINGS_ACCOUNT    Command = "FETCH_LIST_SAVINGS_ACCOUNT"
	CONFIRM_TRANSACTION           Command = "CONFIRM_TRANSACTION"
	REQUEST_SIGNATURE             Command = "REQUEST_SIGNATURE"
)

func (c Command) Equals(cmd Command) bool {
	return strings.Compare(string(c), string(cmd)) == 0
}

func (c Command) ToString() string {
	return string(c)
}

type Status string

const (
	SUCCESS Status = "success"
	ERROR   Status = "error"
)

func (s Status) Equals(st Status) bool {
	return strings.Compare(string(s), string(st)) == 0
}

func (s Status) ToString() string {
	return string(s)
}
