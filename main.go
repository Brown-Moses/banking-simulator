package main

import (
	"codewithwuruem/service"
	"fmt"
)

func main() {

	for {
		fmt.Println("===== ATM SIMULATOR ======")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
		fmt.Println("===========================")

		fmt.Println("Enter your choice: ")
		//call the choice function to handle the user's input and perform the corresponding action
		service.Choice()
	}
}
