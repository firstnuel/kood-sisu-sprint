package sprint

func SplitWhitespaces(s string) []string {
	newStr := []string{}
	start, end := 0, 0

	for i, v := range s {
		if v == ' ' {
			end = i
			if start < end {
				newStr = append(newStr, string(s[start:end]))
			}
			start = i + 1
		}
	}
	if start < len(s) {
		newStr = append(newStr, string(s[start:]))
	}
	return newStr
}
