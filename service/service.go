package service

import (
	"codewithwuruem/model"
	"fmt"
	"os"
)

func Choice(accountbalance *model.Account) {

	var choice int
	var amount float64

	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Your current balance is: ", accountbalance.Balance)
		fmt.Println("---------------------------")

	case 2:
		fmt.Println("Enter the amount to deposit: ")
		fmt.Scan(&amount)
		if amount > 0 {
			accountbalance.Balance += amount
			fmt.Println("Amount successfully deposited:", accountbalance.Balance)
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
			fmt.Println("withdraw Successful: ", accountbalance.Balance)
		}
		fmt.Println("---------------------------")

	case 4:
		fmt.Println("Thanks you for using our ATM.")
		os.Exit(0)

	default:
		fmt.Println("Invalid choice. please try again.")

	}
}
