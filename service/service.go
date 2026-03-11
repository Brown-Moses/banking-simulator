package service

import (
	"codewithwuruem/model"
	"fmt"

	"golang.org/x/text/message"
)

var p = message.NewPrinter(message.MatchLanguage("en"))

// check the balance of the account
func CheckBalance(accountbalance *model.Account) {

	p.Println("Your current balance is: ", accountbalance.Balance)
	fmt.Println("---------------------------")
}

// deposit money into the account
func DepositMoney(accountbalance *model.Account, amount float64) {

	if amount > 0 {
		accountbalance.Balance += amount
		p.Println("Amount successfully deposited:", accountbalance.Balance)
	} else {
		fmt.Println("Invalid AMount")

	}
	fmt.Println("---------------------------")
}

// withdraw money from the account
func WithdrawMoney(accountbalance *model.Account, amount float64) {

	if amount > accountbalance.Balance {
		fmt.Println("insufficient amount")

	} else {
		accountbalance.Balance -= amount
		p.Println("withdraw Successful: ", accountbalance.Balance)
	}
	fmt.Println("---------------------------")
}
