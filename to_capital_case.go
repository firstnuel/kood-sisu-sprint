package sprint

func ToCapitalCase(s string) string {
	var result string
	capitalizeNext := true

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			if capitalizeNext {
				// Convert lowercase letter to uppercase if needed
				if ch >= 'a' && ch <= 'z' {
					ch -= 'a' - 'A'
				}
				capitalizeNext = false
			} else {
				// Convert uppercase letter to lowercase if needed
				if ch >= 'A' && ch <= 'Z' {
					ch += 'a' - 'A'
				}
			}
		} else {
			// If not a letter, reset to capitalize the next letter
			capitalizeNext = true
		}

		result += string(ch)
	}

	return result
}
