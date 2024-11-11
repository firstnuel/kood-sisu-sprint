package sprint

func isPrime(n int) bool {
	num := n
	for n > 2 {
		if num%(n-1) == 0 {
			return false
		}
		n--
	}
	return true
}

func NextPrime(n int) int {
	for {
		if isPrime(n) {
			return n
		}
		n++
	}
}
