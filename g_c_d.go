package sprint

func GCD(a, b int) int {
	n := a
	if a < b {
		n = b
	}

	for i := n; i > 1; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}
	return 1
}
