package isogram
import (
	"strings"
	"unicode"
)

// IsIsogram returns true if a string has no repeating letters.
func IsIsogram(input string) bool {
	var (
		lowerCase  string
		collection map[rune]bool
	)

	collection = map[rune]bool{}
	lowerCase = strings.ToLower(input)

	for _, v := range lowerCase {
		if unicode.IsLetter(v) && collection[v] {
			return false
		}
		collection[v] = true
	}

	return true

}