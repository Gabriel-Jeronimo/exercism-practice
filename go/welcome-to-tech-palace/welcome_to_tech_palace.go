package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return fmt.Sprintf(strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine))
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	oldMsg = strings.Replace(oldMsg, "*", " ", -1)
	oldMsg = strings.Replace(oldMsg, "\n", " ", -1)
	oldMsg = strings.Trim(oldMsg, " ")

	return oldMsg
}
