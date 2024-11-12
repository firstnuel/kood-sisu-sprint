package sprint

func IsNumeric(s string) bool {

	for _, v := range s {
		if !(v >= '0' && v <= '9') {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {

	for _, v := range s {
		if !(v >= 'a' && v <= 'z') {
			return false
		}
	}
	return true
}

func IsUpper(s string) bool {

	for _, v := range s {
		if !(v >= 'A' && v <= 'Z') {
			return false
		}
	}
	return true
}

func IsAlphanumeric(s string) bool {

	for _, v := range s {
		if !(IsUpper(string(v)) || IsNumeric(string(v)) || IsLower(string(v))) {
			return false
		}
	}
	return true
}

func ArrCountIf(f func(string) bool, a []string) int {
	count := 0

	for _, v := range a {
		if f(v) {
			count++
		}
	}
	return count

}
