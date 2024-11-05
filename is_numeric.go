package sprint

func IsNumeric(s string) bool {

	for _, v := range s {
		if !(rune(v) >= 48 && rune(v) <= 57) {
			return false
		}
	}
	return true
}
