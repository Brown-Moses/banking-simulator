package main

import (
	"codewithwuruem/model"
	"codewithwuruem/service"
	"fmt"
)

// Define a global variable to hold the account balance
var account = model.Account{Balance: 0}

func main() {

	if service.PINVerification(&account) == true {

		for {
			fmt.Println("===== ATM SIMULATOR ======")
			fmt.Println("1. Check Balance")
			fmt.Println("2. Deposit Money")
			fmt.Println("3. Withdraw Money")
			fmt.Println("4. Exit")
			fmt.Println("===========================")

			fmt.Println("Enter your choice: ")
			//call the choice function to handle the user's input and perform the corresponding action
			service.Choice(&account)
		}

	} else {
		fmt.Println("PIN verification failed. Exiting the program.")
	}
}
