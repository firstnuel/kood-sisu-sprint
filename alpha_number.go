package sprint

func AlphaNumber(n int) string {
	acc := ""

	if n < 0 {
		n = -n
		acc += "-"
	}

	divisor := 1

	for i := 10; i < n; i *= 10 {
		divisor *= 10
	}

	for divisor > 0 {
		digit := n / divisor
		str := string(97 + digit)
		acc += str
		n = n % divisor
		divisor /= 10
	}

	return acc
}
