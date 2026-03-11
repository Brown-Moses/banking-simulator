package service

import (
	"codewithwuruem/model"
	"fmt"
	"os"

	"golang.org/x/text/message"
)

func Choice(accountbalance *model.Account) {

	var choice int
	var amount float64

	//formating the output to be more user-friendly
	p := message.NewPrinter(message.MatchLanguage("en"))

	fmt.Scan(&choice)

	switch choice {
	case 1:
		p.Println("Your current balance is: ", accountbalance.Balance)
		fmt.Println("---------------------------")

	case 2:
		fmt.Println("Enter the amount to deposit: ")
		fmt.Scan(&amount)
		if amount > 0 {
			accountbalance.Balance += amount
			p.Println("Amount successfully deposited:", accountbalance.Balance)
		} else {
			fmt.Println("Invalid AMount")

		}
		fmt.Println("---------------------------")

	case 3:
		fmt.Println("Enter the amount to withdraw: ")
		fmt.Scan(&amount)
		if amount > accountbalance.Balance {
			fmt.Println("insufficient amount")

		} else {
			accountbalance.Balance -= amount
			p.Println("withdraw Successful: ", accountbalance.Balance)
		}
		fmt.Println("---------------------------")

	case 4:
		fmt.Println("Thanks you for using our ATM.")
		os.Exit(0)

	default:
		fmt.Println("Invalid choice. please try again.")

	}
}
