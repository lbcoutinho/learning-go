package address

import (
	"strings"
)

// Check if address has a valid type and return the type
func AddressType(address string) string {
	validTypes := []string{"street", "road", "alley", "highway"}

	firstWord := strings.Split(strings.ToLower(address), " ")[0]

	isValidType := false
	for _, t := range validTypes {
		if t == firstWord {
			isValidType = true
		}
	}

	if isValidType {
		return firstWord
	}

	return "Invalid type"
}
