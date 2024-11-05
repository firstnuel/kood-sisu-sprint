package sprint

func ToCapitalCase(s string) string {
	var result string
	capitalizeNext := true

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			if capitalizeNext {
				// Capitalize the first letter if it's lowercase
				if ch >= 'a' && ch <= 'z' {
					ch -= 'a' - 'A'
				}
				capitalizeNext = false
			} else {
				// Convert uppercase letters to lowercase for subsequent characters
				if ch >= 'A' && ch <= 'Z' {
					ch += 'a' - 'A'
				}
			}
		} else {
			// Any non-alphabetic character resets the capitalization status for the next letter
			capitalizeNext = true
		}

		result += string(ch)
	}

	return result
}
