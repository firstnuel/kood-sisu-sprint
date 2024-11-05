package sprint

import "strings"

// Helper function to validate the base
func is_ValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	// Check for invalid characters and uniqueness
	seen := make(map[rune]bool)
	for _, ch := range base {
		if ch == '+' || ch == '-' || seen[ch] {
			return false
		}
		seen[ch] = true
	}

	return true
}

// ConvertAnyToDec converts a string `s` in a given `base` to its decimal integer value
func ConvertAnyToDec(s string, base string) int {
	if !is_ValidBase(base) {
		return 0
	}

	baseLen := len(base)
	result := 0

	for _, char := range s {
		index := strings.IndexRune(base, char)
		if index == -1 { // Character not in base
			return 0
		}

		// Accumulate result by shifting and adding the new digit
		result = result*baseLen + index
	}

	return result
}
