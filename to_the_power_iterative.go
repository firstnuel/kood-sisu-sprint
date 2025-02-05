package sprint

func ToThePowerIterative(n int, power int) int {
	pwr := 1

	if power < 0 || n > 20 {
		return 0
	}
	for i := 0; i < power; i++ {
		pwr *= n
	}
	return pwr
}
