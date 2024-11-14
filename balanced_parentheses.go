package sprint

func BalancedParentheses(input string) bool {
	// Stack to keep track of opening parentheses
	var stack []rune
	// Map to define matching pairs
	matches := map[rune]rune{')': '(', ']': '[', '}': '{'}

	// Iterate over each character in the input string
	for _, char := range input {
		switch char {
		// Push opening brackets onto the stack
		case '(', '[', '{':
			stack = append(stack, char)
		// Check for matching closing brackets
		case ')', ']', '}':
			// If stack is empty or top of stack is not a match, return false
			if len(stack) == 0 || stack[len(stack)-1] != matches[char] {
				return false
			}
			// Pop the top of the stack
			stack = stack[:len(stack)-1]
		}
	}

	// If stack is empty, all parentheses are balanced
	return len(stack) == 0
}
