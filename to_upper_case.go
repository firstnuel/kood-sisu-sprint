package sprint

func ToUpperCase(s string) string {
	newStr := ""
	for _, v := range s {
		if rune(v) >= 97 && rune(v) <= 122 {
			newStr += string(rune(v) - 32)
		} else {
			newStr += string(v)
		}
	}
	return newStr
}
