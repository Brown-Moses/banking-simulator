package main

import (
	"codewithwuruem/model"
	"codewithwuruem/service"
	"fmt"
	"os"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {

	// Define a global variable to hold the account balance
	var account = model.Account{Balance: 0, PIN: "1234"}
	var choice int
	var amount float64

	p := message.NewPrinter(language.English)

	if service.VerifyPIN() {

		for {
			fmt.Println("===== ATM SIMULATOR ======")
			fmt.Println("1. Check Balance")
			fmt.Println("2. Deposit Money")
			fmt.Println("3. Withdraw Money")
			fmt.Println("4. Transaction History")
			fmt.Println("5. Exit")
			fmt.Println("===========================")

			fmt.Println("Enter your choice: ")
			fmt.Scan(&choice)

			switch choice {
			//check balance
			case 1:
				service.CheckBalance(&account)
				p.Println("Your current balance is: $", account.Balance)
				fmt.Println("===========================")

				//deposit money
			case 2:
				fmt.Print("Enter the amount to deposit: ")
				fmt.Scan(&amount)
				if err := service.DepositMoney(&account, amount); err != nil {
					fmt.Println("Error:", err)
				} else {
					p.Println("Deposit successful. Balance: $", account.Balance)
				}
				fmt.Println("===========================")

				//withdraw money
			case 3:
				fmt.Print("Enter the amount to withdraw: ")
				fmt.Scan(&amount)
				if err := service.WithdrawMoney(&account, amount); err != nil {
					fmt.Println("Error:", err)
				} else {
					p.Println("Withdrawal successful. Balance: $", account.Balance)
				}
				fmt.Println("===========================")

			case 4:
				fmt.Println("Transaction History:")
				fmt.Println("===========================")
				transactions := service.GetTransactionsHistory(&account)
				for _, transaction := range transactions {
					p.Printf("%s: $%.2f on %s\n", transaction.Type, transaction.Amount, transaction.Timestamp.Format("2006-01-02 15:04:05"))
				}
				fmt.Println("===========================")

			case 5:
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
