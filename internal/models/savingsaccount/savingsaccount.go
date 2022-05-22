package savingsaccount

type SettleType string

const (
	SETTLE_ALL                  SettleType = "SETTLE_ALL"
	ROLLOVER_PRINCIPLE          SettleType = "ROLLOVER_PRINCIPLE"
	ROLLOVER_PRINCIPLE_INTEREST SettleType = "ROLLOVER_PRINCIPLE_INTEREST"
)

type SavingsAccount struct {
	SavingsAccountID     string     `json:"savingsaccount_id,omitempty"`
	ProductTypeName      string     `json:"product_type,omitempty"`
	BankAccountID        string     `json:"bankaccount_id,omitempty"`
	SavingsAmount        float64    `json:"savings_amount,omitempty"`
	InterestAmount       float64    `json:"estimated_interest_amount,omitempty"`
	ActualInterestAmount float64    `json:"actual_interest_amount,omitempty"`
	StartTime            string     `json:"open_time,omitempty"`
	EndTime              string     `json:"settle_time,omitempty"`
	SavingsPeriod        int64      `json:"savings_period,omitempty"`
	SettleInstruction    SettleType `json:"settle_instruction,omitempty"`
	OwnerPhone           string     `json:"owner_phone,omitempty"`
	OwnerID              string     `json:"customer_id,omitempty"`
	InterestRate         float64    `json:"interest_rate,omitempty"`
	CreationConfirmed    string     `json:"creation_confirmed,omitempty"`
	SettleConfirmed      string     `json:"settle_confirmed,omitempty"`
	Currency             string     `json:"currency,omitempty"`
	// 1 for creation successfully, pending creation confirmation
	// 2 for creation confirmed, pending settle
	// 3 for settle successfully, pending settle confirmation
	// 4 for settle confirmed
	ConfirmStatus int64 `json:"confirm_status,omitempty"`
}
