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
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}
