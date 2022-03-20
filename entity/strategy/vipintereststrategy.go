package strategy

type VIPInterestStrategy struct {
	IInterestByCustomerGenreCalculator
}

const additionalVIPInterest float64 = 0.2

func (s VIPInterestStrategy) GetEffectiveInterest(baseInterest float64) (effectiveInterest float64) {
	return baseInterest + additionalVIPInterest
}
