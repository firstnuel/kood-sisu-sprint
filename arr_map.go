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

func IsNegative(n int) bool {
	return n < 0
}

func ArrMap(f func(int) bool, a []int) []bool {
	bSlice := make([]bool, len(a))

	for i, v := range a {
		bSlice[i] = f(v)
	}
	return bSlice
}
