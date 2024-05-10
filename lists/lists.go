package main

import (
	"fmt"

	"example.com/lists/make"
	"example.com/lists/maps"
)

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	var productName [4]string = [4]string{"A Book"}
	productName[1] = "A carpet"
	prices := [4]float64{10.00, 9.99, 45.99, 20.0}
	fmt.Println(prices)
	fmt.Println(productName)

	featuredPrices := prices[1:]
	featuredPrices[0] = 199.99
	fmt.Println(featuredPrices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))

	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	fmt.Println(prices)

	dynamic := []float64{10.09, 18.99}

	fmt.Println(dynamic[0:1])
	dynamic[1] = 9.22
	updatePrices := append(dynamic, 1.44)

	fmt.Println(updatePrices, dynamic)

	discountPrices := []float64{100, 100, 100}
	discountPrices = append(discountPrices, updatePrices...)

	fmt.Println(discountPrices)

	practice()

	maps.Maps()

	make.MakeArray()

}

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

func practice() {
	hobbies := [3]string{"Sports", "Cooking", "Reading"}
	fmt.Println(hobbies)

	// 2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// 3
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	// 4
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// 5
	courseGoals := []string{"Learn Go!", "Learn all the basics"}
	fmt.Println(courseGoals)

	// 6
	courseGoals[1] = "Learn all the details"
	courseGoals = append(courseGoals, "Learn all the basics!")
	fmt.Println(courseGoals)

	// 7
	products := []Product{
		{
			"firstProduct",
			"First product",
			145.00,
		},
		{
			"second product",
			"Second product",
			120.00,
		},
		{
			"third product",
			"Third product",
			12.01,
		},
	}

	newProduct := Product{
		"new product",
		"New product",
		19.01,
	}

	products = append(products, newProduct)
	fmt.Println(products)
}
