package database

import "bankserver/entity/savingsaccount"

func (c DatabaseConnection) SaveCreateNewSavingsAccount(savingsAccount savingsaccount.SavingsAccount) (err error) {
	// Note: ID of type X-ABCD-MNPQ-1234
	// TODO: create an utils to do this, then check existence in db
	sql := `INSERT INTO savingsaccount (
			savingsaccount_id,
			amount,
			period,
			interest_rate,
			interest_amount,
			actual_interest_amount,
			open_time,
			settle_time,
			type,
			creation_confirmed,
			settle_confirmed,
			settle_instruction,
			currency
		)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)
		`
	return c.insert(
		sql,
		savingsAccount.SavingsAccountID,
		savingsAccount.SavingsAmount,
		savingsAccount.SavingsPeriod,
		savingsAccount.InterestRate,
		savingsAccount.InterestAmount,
		0,
		savingsAccount.StartTime,
		savingsAccount.EndTime,
		savingsAccount.ProductTypeName,
		savingsAccount.CreationConfirmed,
		savingsAccount.SettleConfirmed,
		string(savingsAccount.SettleInstruction),
		savingsAccount.Currency,
	)
}

func (c DatabaseConnection) SaveSettleSavingsAccount(savingsAccountID string, settleTime string, actualInterestAmt float64, confirmed string) (err error) {
	sql := `UPDATE savingsaccount
			SET settle_time=?,
				actual_interest_amount=?
			WHERE savingsaccount_id=?`
	return c.update(sql, settleTime, actualInterestAmt, savingsAccountID)
}

func (c DatabaseConnection) SaveSavingAccountCreationConfirmationStatus(savingsAccountID string, confirmed string) (err error) {
	sql := `UPDATE savingsaccount
			SET creation_confirmed=?
			WHERE savingsaccount_id=?`
	return c.update(sql, confirmed, savingsAccountID)
}

func (c DatabaseConnection) SaveSavingAccountSettleConfirmationStatus(savingsAccountID string, confirmed bool) (err error) {
	sql := `UPDATE savingsaccount
			SET settle_confirmed=?
			WHERE savingsaccount_id=?`
	return c.update(sql, confirmed, savingsAccountID)
}
