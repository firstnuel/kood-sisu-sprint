package sprint

func CapitalizeWords(s string) string {
	result := []rune(s)
	newWord := true

	for i := 0; i < len(result); i++ {
		if (result[i] >= 'a' && result[i] <= 'z') || (result[i] >= 'A' && result[i] <= 'Z') || (result[i] >= '0' && result[i] <= '9') {
			if newWord {
				if result[i] >= 'a' && result[i] <= 'z' {
					// Capitalize the first letter of the word
					result[i] -= 'a' - 'A'
				}
				newWord = false
			} else {
				// Convert remaining characters to lowercase if needed
				if result[i] >= 'A' && result[i] <= 'Z' {
					result[i] += 'a' - 'A'
				}
			}
		} else {
			// If the character is not alphanumeric, mark the start of a new word
			newWord = true
		}
	}

	return string(result)
}
