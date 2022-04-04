package savingsaccount

type SettleType string

const (
	SETTLE_ALL                  SettleType = "SETTLE_ALL"
	ROLLOVER_PRINCIPLE          SettleType = "ROLLOVER_PRINCIPLE"
	ROLLOVER_PRINCIPLE_INTEREST SettleType = "ROLLOVER_PRINCIPLE_INTEREST"
)

type SavingsAccount struct {
	SavingsAccountID     string     `json:"savingsaccount_id"`
	ProductTypeName      string     `json:"product_type"`
	BankAccountID        string     `json:"bankaccount_id"`
	SavingsAmount        float64    `json:"savings_amount"`
	InterestAmount       float64    `json:"estimated_interest_amount"`
	ActualInterestAmount float64    `json:"actual_interest_amount"`
	StartTime            string     `json:"open_time"`
	EndTime              string     `json:"settle_time"`
	SavingsPeriod        int64      `json:"savings_period"`
	SettleInstruction    SettleType `json:"settle_instruction"`
	OwnerPhone           string     `json:"owner_phone"`
	OwnerID              string     `json:"customer_id"`
	InterestRate         float64    `json:"interest_rate"`
	CreationConfirmed    bool       `json:"creation_confirmed"`
	SettleConfirmed      bool       `json:"settle_confirmed"`
	Currency             string     `json:"currency"`
}
