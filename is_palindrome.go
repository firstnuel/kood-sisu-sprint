package sprint

func IsPalindrome(s string) bool {
	n := len(s) - 1

	for i := 0; i < len(s); i++ {
		if s[i] != s[n] && n >= 0 {
			return false
		}
		n--
	}
	return true
}
