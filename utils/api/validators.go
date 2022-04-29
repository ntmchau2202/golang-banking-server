package api

import (
	"bankserver/entity/factory"
	"bankserver/entity/savingsaccount"
	"bankserver/entity/savingsproduct"
	"bankserver/utils"
	"errors"
	"fmt"
	"strings"
	"time"
)

func validatePhone(phone string) (err error) {
	if len(phone) == 0 {
		return errors.New("missing customer phone")
	}

	if len(phone) < 10 || len(phone) > 12 {
		return errors.New("invalid phone number")
	}
	return nil
}

func validatePassword(pwd string) (err error) {
	if len(pwd) == 0 {
		return errors.New("missing password")
	}
	return nil
}

func validateBankAccountID(id string) (err error) {
	if len(id) == 0 {
		return errors.New("missing bank account id")
	}
	if len(id) != 4 {
		return errors.New("invalid bank account id")
	}
	_, err = factory.NewBankAccountFactory().GetBankAccountByID(id)
	if err != nil && err.Error() == "no bank account with id given" {
		return
	}
	return nil
}

func validateSavingsType(savingsType string) (err error) {
	if len(savingsType) == 0 {
		return errors.New("missing savings type")
	}
	fmt.Println(savingsproduct.SavingsProductTypeName)
	fmt.Println(savingsType)
	for _, item := range savingsproduct.SavingsProductTypeName {
		if item == savingsType {
			fmt.Println("Gotcha")
			return nil
		}
	}
	return errors.New("no such savings product")
}

func validateSavingsPeriod(savingsType string, savingsPeriod int) (err error) {
	if err = validateSavingsType(savingsType); err != nil {
		return
	}

	if savingsPeriod <= 0 {
		return errors.New("invalid savings period")
	}

	fmt.Println("Interestrate map:", savingsproduct.SavingsProductType[savingsType].InterestRate)
	for period := range savingsproduct.SavingsProductType[savingsType].InterestRate {
		fmt.Println("period:", period)
		fmt.Println("input savingsPeriod:", savingsPeriod)
		if period == savingsPeriod {
			return nil
		}
	}
	return errors.New("no savings period for given savings product")
}

func validateAmount(amount float64) (err error) {
	if amount <= 0 {
		return errors.New("invalid amount")
	}
	return nil
}

func validateSettleInstruction(instruction string) (err error) {
	instructionU := strings.ToUpper(instruction)
	if strings.Compare(instructionU, string(savingsaccount.SETTLE_ALL)) != 0 &&
		strings.Compare(instructionU, string(savingsaccount.ROLLOVER_PRINCIPLE)) != 0 &&
		strings.Compare(instructionU, string(savingsaccount.ROLLOVER_PRINCIPLE_INTEREST)) != 0 {
		return errors.New("invalid settle instruction")
	}
	return nil
}

func validateCurrency(currency string) (err error) {
	currencyU := strings.ToUpper(currency)
	if strings.Compare(currencyU, "VND") != 0 &&
		strings.Compare(currencyU, "USD") != 0 {
		// currently only support these 2 currency
		return errors.New("invalid currency unit")
	}
	return nil
}

func validateTime(t string) (err error) {
	_, err = time.Parse(time.RFC1123, t)
	if err != nil {
		fmt.Println(err)
		return errors.New("invalid time format")
	}
	return nil
}

func validateInterestRate(rate float64) (err error) {
	if rate <= 0 {
		return errors.New("invalid interest rate")
	}
	return nil
}

func validateSavingsAccountID(id string) (err error) {
	chunk := strings.Split(id, "-")
	if len(chunk) != 4 {
		return errors.New("invalid savings account id")
	}

	if len(chunk[0]) != 1 {
		return errors.New("invalid savings account id")
	}

	if utils.HasDigit(chunk[0]) || utils.HasDigit(chunk[1]) || utils.HasDigit(chunk[2]) {
		return errors.New("invalid savings account id")
	}

	if utils.HasLetter(chunk[3]) {
		return errors.New("invalid savings account id")
	}

	if _, err = factory.NewSavingsAccountFactory().GetSavingsAccountByID(id); err != nil && err.Error() == "no savings account with id given" {
		return err
	}

	return nil
}
