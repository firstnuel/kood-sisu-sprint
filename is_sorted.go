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

func IsSorted(f func(a, b string) int, arr []string) bool {

	for i := 0; i < len(arr)-1; i++ {
		if f(arr[i], arr[i+1]) < 0 {
			return false
		}
	}
	return true
}
