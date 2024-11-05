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
