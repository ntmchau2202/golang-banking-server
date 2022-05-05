package strategy

type NormalInterestStrategy struct {
	IInterestByCustomerGenreCalculator
}

const additionalNormalStrategy float64 = 0.2

func (s NormalInterestStrategy) GetEffectiveInterest(baseInterest float64) (effectiveInterest float64) {
	return baseInterest
}
