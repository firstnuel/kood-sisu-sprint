package sprint

func GetLastRune(s string) rune {
	n := len(s) - 1
	return rune(s[n])
}
