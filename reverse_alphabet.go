package sprint

func ReverseAlphabet(step int) string {
	acc := ""
	for i := 122; i >= 97; i -= step {
		acc += string(i)
	}
	return acc
}
