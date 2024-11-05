package sprint

func ToCapitalCase(s string) string {
	var result []rune
	isFirstChar := true
	for _, char := range s {
		if isFirstChar && isAlphanumeric(char) {
			result = append(result, toUpper(char))
			isFirstChar = false
		} else if char == ' ' {
			result = append(result, char)
			isFirstChar = true
		} else {
			result = append(result, toLower(char))
		}
	}
	return string(result)
}

func isAlphanumeric(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || ('0' <= r && r <= '9')
}

func toUpper(r rune) rune {
	if 'a' <= r && r <= 'z' {
		return r - 'a' + 'A'
	}
	return r
}

func toLower(r rune) rune {
	if 'A' <= r && r <= 'Z' {
		return r - 'A' + 'a'
	}
	return r
}
