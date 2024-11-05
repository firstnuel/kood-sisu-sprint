package sprint

import (
	"strings"
)

// Helper function to validate the base
func isValid_Base(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, ch := range base {
		if ch == '+' || ch == '-' || seen[ch] {
			return false
		}
		seen[ch] = true
	}
	return true
}

// Convert a number in any base to decimal
func baseToDecimal(nbr, baseFrom string) int {
	baseLen := len(baseFrom)
	result := 0

	for _, ch := range nbr {
		index := strings.IndexRune(baseFrom, ch)
		if index == -1 { // Character not in baseFrom
			return -1
		}
		result = result*baseLen + index
	}

	return result
}

// Convert a decimal number to any base
func decimalToBase(num int, baseTo string) string {
	if num == 0 {
		return string(baseTo[0])
	}
	baseLen := len(baseTo)
	result := ""

	for num > 0 {
		remainder := num % baseLen
		result = string(baseTo[remainder]) + result
		num /= baseLen
	}

	return result
}

// ConvertAnyToAny converts a number `nbr` from `baseFrom` to `baseTo`
func ConvertAnyToAny(nbr, baseFrom, baseTo string) string {
	// Validate bases
	if !isValid_Base(baseFrom) || !isValid_Base(baseTo) {
		return "0"
	}

	// Convert from baseFrom to decimal
	decimalValue := baseToDecimal(nbr, baseFrom)
	if decimalValue == -1 {
		return "0" // Invalid character in `nbr` for baseFrom
	}

	// Convert from decimal to baseTo
	return decimalToBase(decimalValue, baseTo)
}
