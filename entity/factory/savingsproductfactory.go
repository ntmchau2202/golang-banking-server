package factory

import (
	"bankserver/database"
	"bankserver/entity/savingsproduct"
	"errors"
)

type savingsProductFactory struct {
}

func NewSavingsProductFactory() *savingsProductFactory {
	return &savingsProductFactory{}
}

func isProductTypeExist(name string) bool {
	_, exist := savingsproduct.SavingsProductType[name]
	return exist
}

func (f savingsProductFactory) PutProductType(name string, product savingsproduct.SavingsProduct) {
	if !isProductTypeExist(name) {
		savingsproduct.SavingsProductType[name] = product
	}
}

func (f savingsProductFactory) FetchAllSavingsProduct() (err error) {
	for _, item := range savingsproduct.SavingsProductTypeName {
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

func (f savingsProductFactory) GetSavingsProductByName(name string) (product savingsproduct.SavingsProduct, err error) {
	if product, exist := savingsproduct.SavingsProductType[name]; exist {
		return product, nil
	} else {
		return product, errors.New("product not found")
	}
}
