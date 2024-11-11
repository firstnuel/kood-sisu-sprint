package sprint

func FactorialIterative(n int) int {
	fact := 1

	if n == 0 {
		return 1
	}

	if n < 0 || n > 20 {
		return 0
	}

	for i := n; i >= 1; i-- {
		fact *= i
	}
	return fact
}
