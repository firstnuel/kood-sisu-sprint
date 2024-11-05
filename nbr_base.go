package sprint

// Helper function to check if the base is valid
func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	// Check for '+' or '-' in base
	for _, ch := range base {
		if ch == '+' || ch == '-' {
			return false
		}
	}

	// Check for duplicates
	seen := make(map[rune]bool)
	for _, ch := range base {
		if seen[ch] {
			return false
		}
		seen[ch] = true
	}

	return true
}

func NbrBase(n int, base string) string {
	if !isValidBase(base) {
		return "NV"
	}

	baseLen := len(base)
	isNegative := n < 0
	if isNegative {
		n = -n
	}

	// Handle case where n is zero
	if n == 0 {
		return string(base[0])
	}

	result := ""
	for n > 0 {
		remainder := n % baseLen
		result = string(base[remainder]) + result
		n /= baseLen
	}

	if isNegative {
		result = "-" + result
	}

	return result
}
