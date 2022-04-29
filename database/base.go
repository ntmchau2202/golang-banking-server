package database

import (
	"bankserver/entity/customer"
	"bankserver/entity/savingsproduct"
	"fmt"
	"strings"
	"sync"

	"zombiezen.com/go/sqlite"
	db "zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

// singleton
const (
	dbPath string = "./resource/bankDB.db"
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
			fmt.Println(err)
			return
		}
		databaseConnection.dbConn = dBConnection
		initConfigs()
	})
	return databaseConnection, nil
}

func initConfigs() (err error) {
	savingsproduct.SavingsProductType = make(map[string]savingsproduct.SavingsProduct)
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
			product, err := databaseConnection.GetSavingsProductDetails(value)
			if err != nil {
				continue
			}
			savingsproduct.SavingsProductType[value] = product
		}
	}

	return
}

func (c DatabaseConnection) insert(stm string, args ...interface{}) (err error) {
	execArgs := sqlitex.ExecOptions{}
	execArgs.Args = append(execArgs.Args, args...)
	return sqlitex.Execute(c.dbConn, stm, &execArgs)
}

func (c DatabaseConnection) update(stm string, args ...interface{}) (err error) {
	execArgs := sqlitex.ExecOptions{}
	execArgs.Args = append(execArgs.Args, args...)
	return sqlitex.Execute(c.dbConn, stm, &execArgs)
}
