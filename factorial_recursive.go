package sprint

func FactorialRecursive(n int) int {
	if n == 0 {
		return 1
	}

	if n < 0 || n > 20 {
		return 0
	}

	return n * FactorialRecursive(n-1)
}
