package strategy

import (
	"bankserver/entity/customer"
	"bankserver/entity/savingsproduct"
	"reflect"
)

type InterestStrategyFactory struct {
}

var factory InterestStrategyFactory

func GetFactory() InterestStrategyFactory {
	return factory
}

func (f InterestStrategyFactory) GetStrategyByProduct(product savingsproduct.SavingsProduct, period int) IInterestByProductCalculator {
	if reflect.TypeOf(product).Name() == "TraditionalOnlineSavingsProduct" {
		return OnlineInterestStrategy{
			Product: product,
			Period:  period,
		}
	}
	return nil
}

func (f InterestStrategyFactory) GetStrategyByCustomerGenre(cust customer.Customer) IInterestByCustomerGenreCalculator {
	if reflect.TypeOf(cust).Name() == "NormalCustomer" {
		return NormalInterestStrategy{}
	}
	return nil
}
