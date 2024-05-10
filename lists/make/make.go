package make

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func MakeArray() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Make")
	userNames := make([]string, 2, 10)

	userNames[0] = "Julie"
	userNames[1] = "Xean"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)
	courseRatings["go"] = 4.9
	courseRatings["react"] = 4.0
	courseRatings["angular"] = 4.1
	courseRatings["node"] = 4.7

	courseRatings.output()

	for index, value := range userNames {
		fmt.Println("Index: ", index)
		fmt.Println("Value: ", value)
	}

	for key, value := range courseRatings {
		fmt.Println("key: ", key)
		fmt.Println("Value: ", value)
	}
}
