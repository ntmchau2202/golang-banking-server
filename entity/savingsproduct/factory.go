package savingsproduct

type SavingsProductFactory struct {
}

var SavingsProductType []SavingsProduct

func isProductTypeExist(product SavingsProduct) bool {
	for _, item := range SavingsProductType {
		if item.ProductName == product.ProductName {
			return true
		}
	}
	return false
}

func (f SavingsProductFactory) PutProductType(product SavingsProduct) {
	if !isProductTypeExist(product) {
		SavingsProductType = append(SavingsProductType, product)
	}
}

func (f SavingsProductFactory) FetchAllSavingsProduct() (err error) {
	for _, item := 
}
