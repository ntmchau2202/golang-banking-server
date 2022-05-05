package strategy

import "core-banking-server/internal/models/savingsproduct"

type OnlineInterestStrategy struct {
	IInterestByProductCalculator
	Product savingsproduct.SavingsProduct
	Period  int
}

const additionalOnlineInterest float64 = 0.3

func (s OnlineInterestStrategy) GetEffectiveInterest() (effectiveInterest float64) {
	return s.Product.GetBaseInterestRateOfMonth(s.Period) + additionalOnlineInterest
}
