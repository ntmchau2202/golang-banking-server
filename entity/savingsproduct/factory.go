package savingsproduct

import (
	"bankserver/database"
	"errors"
)

type SavingsProductFactory struct {
}

var SavingsProductType map[string]SavingsProduct
var SavingsProductTypeName []string

func isProductTypeExist(name string) bool {
	_, exist := SavingsProductType[name]
	return exist
}

func (f SavingsProductFactory) PutProductType(name string, product SavingsProduct) {
	if !isProductTypeExist(name) {
		SavingsProductType[name] = product
	}
}

func (f SavingsProductFactory) FetchAllSavingsProduct() (err error) {
	for _, item := range SavingsProductTypeName {
		db, err := database.GetDBConnection()
		if err != nil {
			return err
		}

		product, err := db.GetSavingsProductDetails(item)
		if err != nil {
			return err
		}

		f.PutProductType(item, product)
	}
	return nil
}

func (f SavingsProductFactory) GetSavingsProductByName(name string) (product SavingsProduct, err error) {
	if product, exist := SavingsProductType[name]; exist {
		return product, nil
	} else {
		return product, errors.New("product not found")
	}
}
