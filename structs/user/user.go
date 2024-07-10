package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	Birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdminUser(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "admin",
			lastName:  "admin",
			Birthdate: "---",
			createdAt: time.Now(),
		},
	}
}

// (u user) -> receiver
func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.Birthdate)
}

func (u *User) ClearUsername() {

	u.firstName = ""
	u.lastName = ""
}

func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("Required fields must be filled")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		Birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
