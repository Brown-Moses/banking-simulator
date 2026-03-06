package service

import (
	"fmt"
	"os"
)

func Choice() {
	var balance float64 = 0
	var choice int
	var amount float64

	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Your current balance is: ", balance)
		fmt.Println("---------------------------")

	case 2:
		fmt.Println("Enter the amount to deposit: ")
		fmt.Scan(&amount)
		if amount > 0 {
			balance += amount
			fmt.Println("Amount successfully deposited:", balance)
		} else {
			fmt.Println("Invalid AMount")

		}
		fmt.Println("---------------------------")

	case 3:
		fmt.Println("Enter the amount to withdraw: ")
		fmt.Scan(&amount)
		if amount > balance {
			fmt.Println("insufficient amount")

		} else {
			balance -= amount
			fmt.Println("withdraw Successful: ", amount)
		}
		fmt.Println("---------------------------")

	case 4:
		fmt.Println("Thanks you for using our ATM.")
		os.Exit(0)

	default:
		fmt.Println("Invalid choice. please try again.")

	}
}
