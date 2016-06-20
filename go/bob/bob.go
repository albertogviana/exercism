package bob // package name must match the package name in bob_test.go
import (
	"strings"
	"unicode"
)

const testVersion = 2 // same as targetTestVersion

// Hey return the answer for the question
func Hey(question string) string {
	message := strings.TrimSpace(question)

	switch {
	case message == "":
		return "Fine. Be that way!"
	case CheckRune(message, unicode.IsUpper) && !CheckRune(message, unicode.IsLower):
		return "Whoa, chill out!"
	case message[len(message)-1:] == "?":
		return "Sure."
	default:
		return "Whatever."
	}
}

// CheckRune if the question has symbols
func CheckRune(items string, test func(rune) bool) bool {
	for _, item := range items {
		if test(item) {
			return true
		}
	}
	return false
}
