package sprint

func IsLeapYear(year int) bool {
	if year%100 == 0 && year%400 == 0 {
		return true
	}
	return false
}
