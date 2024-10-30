package sprint

func Accumulate(n int) int {
	sum := 0
	if n < 0 {
		return sum
	}
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}
