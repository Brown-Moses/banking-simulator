package service

import (
	"codewithwuruem/model"
	"errors"
)

// check the balance of the account
func CheckBalance(accountbalance *model.Account) float64 {

	return accountbalance.Balance

}

// deposit money into the account
func DepositMoney(accountbalance *model.Account, amount float64) error {

	if amount <= 0 {
		return errors.New("Invalid AMount")

	}
	accountbalance.Balance += amount
	return nil
}

// withdraw money from the account
func WithdrawMoney(accountbalance *model.Account, amount float64) error {

	if amount <= 0 {
		return errors.New("Invalid Withdrawal Amount")

	}

	if amount > accountbalance.Balance {
		return errors.New("insufficient amount")

	}
	accountbalance.Balance -= amount
	return nil
}
