package main

import (
	"fmt"

	"example.com/structs/user"
)

type customString string

func (text customString) log() {
	fmt.Println(text)
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser, err = user.NewUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "test123")
	admin.OutputUserDetails2()
	admin.ClearUserName()
	admin.OutputUserDetails2()

	user.OutputUserDetails(appUser)
	appUser.ClearUserName()
	appUser.OutputUserDetails2()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
