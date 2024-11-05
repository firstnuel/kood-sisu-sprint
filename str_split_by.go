package sprint

func StrSplitBy(s, sep string) []string {
	newStr := []string{}
	start, end := 0, 0
	n := len(sep)

	for i := 0; i < len(s)-n; i += 1 {
		if s[i:i+n] == sep {
			end = i
			if start < end {
				newStr = append(newStr, s[start:end])
			}
			start = i + 3
		}

	}
	if start < len(s) {
		newStr = append(newStr, s[start:])
	}
	return newStr
}
