package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	var welcomeMsg string = ("Welcome to the Tech Palace," + " " + strings.ToUpper(customer))
	return welcomeMsg
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var border string = (strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine))
	return border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	var cleanStars string = strings.ReplaceAll(oldMsg, "*", "")
	var cleanSpace string = strings.TrimSpace(cleanStars)
	return cleanSpace
}
