package main

import "fmt"

func main() {
	sum := sumup(1, 10, 15, -5)
	fmt.Println(sum)

	numbers := []int{1, 10, 15}
	anotherSum := sumup(1, numbers...)

	fmt.Println(anotherSum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
