package service

import (
	"fmt"
)

func VerifyPIN() bool {

	correctPIN := "1234"
	var inputPIN string

	fmt.Println("Enter your PIN: ")
	fmt.Scan(&inputPIN)
	fmt.Println("---------------------------")

	if inputPIN == correctPIN {
		return true
	} else {
		fmt.Println("Invalid PIN. Please try again.")
		return false
	}
}
