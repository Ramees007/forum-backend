package entity

import (
	"regexp"
)

var regexpEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

//User model
type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

//IsValidForReg - checks if user is valid
func (a *User) IsValidForReg() (err string) {
	// check if the name empty
	if a.Name == "" {
		return "The name is required!"
	}
	// check the name field is between 3 to 120 chars
	if len(a.Name) < 2 || len(a.Name) > 40 {
		return "The name field must be between 2-40 chars!"
	}
	if a.Email == "" {
		return "The email field is required!"
	}

	if a.Password == "" {
		return "The password field is required!"
	}

	if !regexpEmail.MatchString(a.Email) {
		return "The email field should be a valid email address!"
	}
	return ""
}

//IsValidForLogin - checks if user is valid
func (a *User) IsValidForLogin() (err string) {

	if a.Email == "" {
		return "The email field is required!"
	}

	if a.Password == "" {
		return "The password field is required!"
	}

	return ""
}
