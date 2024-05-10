package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountFile)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Welcome to Go bank")
	fmt.Println("Reach us 24/7 ", randomdata.PhoneNumber())

	// for i := 0; i < 2; i++ {
	for {
		presetOption()

		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Your balance is ", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount ")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("New balance ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountFile)
		case 3:
			fmt.Println("withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount ")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("More than your current balance ")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("New balance ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountFile)
		case 4:
			fmt.Println("Good bye")
			fmt.Println("Thank you for choosing our bank")
			return

		}
	}

}
