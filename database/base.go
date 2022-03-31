package database

import (
	"bankserver/entity/customer"
	"bankserver/entity/savingsproduct"
	"strings"
	"sync"

	"zombiezen.com/go/sqlite"
	db "zombiezen.com/go/sqlite"
)

// singleton
const (
	dbPath string = ""
)

type DatabaseConnection struct {
	dbConn *db.Conn
}

var databaseConnection DatabaseConnection

var singletonDB sync.Once

func GetDBConnection() (conn DatabaseConnection, err error) {
	singletonDB.Do(func() {
		dBConnection, err := db.OpenConn(dbPath, sqlite.OpenReadWrite)
		if err != nil {
			return
		}
		databaseConnection.dbConn = dBConnection
		initConfigs()
	})
	return databaseConnection, nil
}

func initConfigs() (err error) {
	stm := databaseConnection.dbConn.Prep("SELECT * FROM configs")
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
			customer.CustomerType = append(customer.CustomerType, value)
		} else if strings.Compare(field, "savings_product_type") == 0 {
			savingsproduct.SavingsProductTypeName = append(savingsproduct.SavingsProductTypeName, value)
		}
	}
	return
}
