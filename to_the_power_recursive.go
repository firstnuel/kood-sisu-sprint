package sprint

func ToThePowerRecursive(n int, power int) int {
	if power < 0 || n > 20 {
		return 0
	}
	if power == 0 {
		return 1
	}

	return n * ToThePowerRecursive(n, power-1)
}
