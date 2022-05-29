package database

import "core-banking-server/internal/models/savingsaccount"

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
			currency,
			status
		)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)
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
		1,
	)
}

func (c DatabaseConnection) SaveSettleSavingsAccount(savingsAccountID string, settleTime string, actualInterestAmt float64, confirmed string) (err error) {
	sql := `UPDATE savingsaccount
			SET settle_time=?,
				actual_interest_amount=?,
				status=?
			WHERE savingsaccount_id=?`
	return c.update(sql, settleTime, actualInterestAmt, 3, savingsAccountID)
}

func (c DatabaseConnection) UpdateAccountBalance(bankAccountID string, newBalance float64) (err error) {
	sql := `UPDATE bankaccount
			SET bankaccount_balance=?
			WHERE bankaccount_id=?`
	return c.update(sql, newBalance, bankAccountID)
}

func (c DatabaseConnection) SaveSavingAccountCreationConfirmationStatus(savingsAccountID string, txnHash string, ipfsHash string) (err error) {
	sql := `UPDATE savingsaccount
			SET creation_confirmed=?,
				status=?,
				open_ipfs_receipt_hash=?
			WHERE savingsaccount_id=?`
	return c.update(sql, txnHash, 2, ipfsHash, savingsAccountID)
}

func (c DatabaseConnection) SaveSavingAccountSettleConfirmationStatus(savingsAccountID string, txnHash string, ipfsHash string) (err error) {
	sql := `UPDATE savingsaccount
			SET settle_confirmed=?,
				status=?,
				settle_ipfs_receipt_hash=?
			WHERE savingsaccount_id=?`
	return c.update(sql, txnHash, 4, ipfsHash, savingsAccountID)
}

func (c DatabaseConnection) AddSavingsAccountToBankAccount(savingsAccountID, bankAccountID string) (err error) {
	sql := `INSERT INTO bankaccount_savingsaccount (bankaccount_id, savingsaccount_id)
			VALUES (?, ?)`
	return c.insert(sql, bankAccountID, savingsAccountID)
}

func (c DatabaseConnection) AddCustomerPublicEncryptionKey(customerPhone, publicKey string) (err error) {
	sql := `UPDATE customer
			SET public_key=?
			WHERE customer_id=?`
	return c.update(sql, publicKey, customerPhone)
}

func (c DatabaseConnection) RemoveCustomerPublicEncryptionKey(customerPhone string) (err error) {
	sql := `UPDATE customer
			SET public_key=""
			WHERE customer_id=?`
	return c.update(sql, customerPhone)
}
