package savingsaccount

import (
	"bankserver/entity/customer"
	"bankserver/entity/savingsproduct"
)

type SettleType string

const (
	SETTLE_ALL                  SettleType = "SETTLE_ALL"
	ROLLOVER_PRINCIPLE          SettleType = "ROLLOVER_PRINCIPLE"
	ROLLOVER_PRINCIPLE_INTEREST SettleType = "ROLLOVER_PRINCIPLE_INTEREST"
)

type SavingsAccount struct {
	SavingsAccountID    string                        `json:"savingsaccount_id"`
	ProductType         savingsproduct.SavingsProduct `json:"product_type"`
	SavingsAmount       float64                       `json:"savings_amount"`
	InterestAmount      float64                       `json:"actual_interest_amount"`
	StartTime           string                        `json:"start_time"`
	EndTime             string                        `json:"end_time"`
	SavingsPeriod       int64                         `json:"savings_period"`
	SettleInstruction   SettleType                    `json:"settle_instruction"`
	Owner               customer.Customer             `json:"customer"`
	InterestRate        float64                       `json:"interest_rate"`
	BlockchainConfirmed bool                          `json:"blockchain_confirmed"`
	Currency            string                        `json:"currency"`
}
