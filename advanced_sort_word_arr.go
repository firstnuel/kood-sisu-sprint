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
	// Bubble sorting
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if f(a[j], a[j+1]) > 0 {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}
