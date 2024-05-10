package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err1 := getUserInput("Revenue: ")

	if err1 != nil {
		fmt.Println(err1)
		return
	}

	expense, err2 := getUserInput("Expense: ")

	if err2 != nil {
		fmt.Println(err2)
		return
	}

	taxRate, err3 := getUserInput("Tax Rate: ")

	if err3 != nil {
		fmt.Println(err3)
		return
	}

	ebt, profit, ratio := calculateFinanicials(revenue, expense, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.1f\n", ratio)

	storeResult(ebt, profit, ratio)

}

func storeResult(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT:%.1f\nProfit:%.1f\nRatio:%.3f", ebt, profit, ratio)
	os.WriteFile("result.txt", []byte(results), 0644)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput < 0 {
		return 0, errors.New("value must be greater than 0")
	}
	return userInput, nil
}

func calculateFinanicials(revenue, expense, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expense
	profit := float64(ebt) * (1 - taxRate/100)
	ratio := float64(ebt) / profit

	return ebt, profit, ratio

}
