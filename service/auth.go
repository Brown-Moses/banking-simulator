package service

import (
	"codewithwuruem/model"
	"fmt"
)

func PINVerification(pin *model.Account) bool {

	correctPIN := "1234"

	fmt.Println("Enter your PIN: ")
	fmt.Scan(&pin.PIN)
	fmt.Println("---------------------------")

	if pin.PIN == correctPIN {
		return true
	} else {
		fmt.Println("Invalid PIN. Please try again.")
		return false
	}
}
