package helper

import "strings"

// variabel with capital leter bicouse is export to main.go
var MyVar = "somthing"

// function with capital leter bicouse is export to main.go
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isVaidticketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isVaidticketNumber
}
