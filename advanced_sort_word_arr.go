package sprint

func StrCompare(a, b string) int {
	if a == b {
		return 0
	}
	n := len(b)
	for i := 0; i < len(a)-n; i++ {
		if a[i:i+n] == b {
			if i == 0 {
				return 1
			}
		}
	}
	return -1
}

func AdvancedSortWordArr(a []string, f func(a, b string) int) []string {
	n := len(a)
	for {
		switched := false
		for i := 0; i < n-1; i++ {
			if f(a[i], a[i+1]) < 1 {
				a[i], a[i+1] = a[i+1], a[i]
				switched = true
			}
		}
		if !switched {
			break
		}
		n--
	}
	return a
}
