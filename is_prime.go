package sprint

func IsPrime(n int) bool {
	num := n
	for n > 2 {
		if num%(n-1) == 0 {
			return false
		}
		n--
	}
	return true
}
