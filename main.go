package main

import (
	"codewithwuruem/model"
	"codewithwuruem/service"
	"fmt"
	"os"
)

func main() {

	// Define a global variable to hold the account balance
	var account = model.Account{Balance: 0, PIN: "1234"}
	var choice int
	var amount float64

	if service.VerifyPIN() {

		for {
			fmt.Println("===== ATM SIMULATOR ======")
			fmt.Println("1. Check Balance")
			fmt.Println("2. Deposit Money")
			fmt.Println("3. Withdraw Money")
			fmt.Println("4. Exit")
			fmt.Println("===========================")

			fmt.Println("Enter your choice: ")
			fmt.Scan(&choice)
			//call the choice function to handle the user's input and perform the corresponding action
			//service.Choice(&account)

			switch choice {
			case 1:
				service.CheckBalance(&account)

			case 2:
				fmt.Print("Enter the amount to deposit: ")
				fmt.Scan(&amount)
				service.DepositMoney(&account, amount)

			case 3:
				fmt.Print("Enter the amount to withdraw: ")
				fmt.Scan(&amount)
				service.WithdrawMoney(&account, amount)

			case 4:
				fmt.Println("Thanks you for using our ATM.")
				os.Exit(0)
			default:
				fmt.Println("Invalid choice. please try again.")
			}
		}

	} else {
		fmt.Println("PIN verification failed. Exiting the program.")
	}
}
