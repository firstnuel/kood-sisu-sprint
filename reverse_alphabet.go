package sprint

func ReverseAlphabet(step int) string {
	acc := ""
	if step <= 0 {
		step = 1
	}
	for i := 122; i >= 97; i -= step {
		acc += string(i)
	}
	return acc
}
