package sprint

func ToCapitalCase(s string) string {
	var result string
	capitalizeNext := true

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			if capitalizeNext {
				// Convert lowercase to uppercase if needed for the first letter
				if ch >= 'a' && ch <= 'z' {
					ch -= 'a' - 'A'
				}
				capitalizeNext = false
			} else {
				// Convert uppercase to lowercase for subsequent letters
				if ch >= 'A' && ch <= 'Z' {
					ch += 'a' - 'A'
				}
			}
		} else {
			// Reset on non-letters, so next letter is capitalized
			capitalizeNext = true
		}

		result += string(ch)
	}

	return result
}
