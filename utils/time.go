package utils

import "time"

func GetCurrentTimeFormatted() (t string) {
	return time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST") // RFC123
}
