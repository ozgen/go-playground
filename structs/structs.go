package main

import (
	"fmt"
	"os"
	"structs/user"
)

type str string

func (txt str) log() {
	fmt.Println(txt)
}
func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	appUser.OutputUserDetails()
	appUser.ClearUsername()
	appUser.OutputUserDetails()
	fmt.Println(appUser.Birthdate) // you can only reach Birthdate because it starts with CAPITAL

	adminUser := user.NewAdminUser("test@test.com", "test123")

	adminUser.OutputUserDetails() // if you add anonymous variable like user without specific name
	// you can also use user method without adminUser.User.OutputUserDetails()
	adminUser.ClearUsername()
	adminUser.OutputUserDetails()

	//in order to overrite string type adding a new method
	// you can use like that
	var test str = "test message"
	test.log()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
