package sprint

func SimpleStrToInt(s string) int {

	acc := 0
	for i := 0; i <= len(s)-1; i++ {
		if acc != 0 || s[i] != '0' {
			digit := int(s[i] - '0')
			acc = acc*10 + digit
		}
	}
	return acc
}
