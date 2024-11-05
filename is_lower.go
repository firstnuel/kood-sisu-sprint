package sprint

func IsLower(s string) bool {

	for _, v := range s {
		if !(rune(v) >= 97 && rune(v) <= 122) {
			return false
		}
	}
	return true
}
