package sprint

func titleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	if rune(s[0]) >= 97 && rune(s[0]) <= 122 {
		return string(rune(s[0])-32) + s[1:]
	}
	return s
}

func ToCapitalCase(s string) string {
	newStr := ""
	n := 0

	for i := 0; i < len(s); i++ {
		if !((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z')) {
			if n < i {
				word := s[n:i]
				newStr += titleCase(word)
			}
			newStr += string(s[i])
			n = i + 1
		}
	}
	if n < len(s) {
		word := s[n:]
		newStr += titleCase(word)
	}
	return newStr
}
