package database

import (
	"strings"
	"sync"

	db "zombiezen.com/go/sqlite"
)

// singleton

var (
	DBConnection *db.Conn
)
var singletonDB sync.Once

func GetDBConnection() (conn *db.Conn, err error) {
	singletonDB.Do(func() {
		// initialize dabtase here
	})
	return
}

func initConfigs() (err error) {
	stm, err := DBConnection.Prepare("SELECT * FROM configs")
	if err != nil {
		return
	}
	defer stm.Finalize()
	for {
		if hasRow, err := stm.Step(); err != nil {
			// handle error
		} else if !hasRow {
			break
		}
		field := stm.GetText("field")
		value := stm.GetText("value")
		if strings.Compare(field, "customer_type") == 0 {
			// call customer factory to put customer type here
		} else if strings.Compare(field, "savings_product_type") == 0 {
			// call savings product factory to put savings product type here
		}
	}
	return
}


