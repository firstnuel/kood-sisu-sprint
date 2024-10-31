package sprint

func Countdown(n int) string {
	acc := ""
	for i := n; i >= 0; i -= 2 {
		acc += string(i+'0') + ", "
	}
	return acc + "0!"
}
