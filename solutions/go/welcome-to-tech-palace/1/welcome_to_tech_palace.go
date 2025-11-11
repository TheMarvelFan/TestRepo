package techpalace

import "fmt"
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	ret := ""
    copyNumsPerLine := numStarsPerLine

    for numStarsPerLine > 0 {
        numStarsPerLine -= 1
        ret += "*"
    }

    ret += fmt.Sprintf("\n%s\n", welcomeMsg)

    for copyNumsPerLine > 0 {
        copyNumsPerLine -= 1
        ret += "*"
    }

    return ret
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	msg := strings.ReplaceAll(oldMsg, "*", "")
    return strings.TrimSpace(msg)
}
