package sprint

func titleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	if s[0] >= 97 && s[0] <= 122 {
		return string(s[0]-32) + s[1:]
	}
	return s
}

func ToCapitalCase(s string) string {
	newStr := ""
	n := 0

	for i := 0; i < len(s); i++ {
		// Check for space or hyphen to signal the end of a word
		if s[i] == ' ' || s[i] == '-' {
			if n < i {
				word := s[n:i]
				newStr += titleCase(word)
			}
			// Append the non-letter separator (space or hyphen)
			newStr += string(s[i])
			n = i + 1 // Move start of next word to the character after the separator
		}
	}
	// Handle the last word if there’s any remaining characters
	if n < len(s) {
		word := s[n:]
		newStr += titleCase(word)
	}

	return newStr
}
