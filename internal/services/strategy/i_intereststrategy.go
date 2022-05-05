package strategy

type IInterestByCustomerGenreCalculator interface {
	GetEffectiveInterest(baseInterest float64) (effectiveInterest float64)
}

type IInterestByProductCalculator interface {
	GetEffectiveInterest() (effectiveInterest float64)
}
