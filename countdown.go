package sprint

func Countdown(n int) string {
	if n <= 0 {
		return "0!"
	}
	acc := ""
	for i := n; i >= 0; i -= 2 {
		acc += string(i + '0')
		if i != 0 {
			acc += ", "
		}
	}
	if n%2 != 0 {
		return acc + "0!"
	}
	return acc + "!"
}
