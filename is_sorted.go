package sprint

func StrCompare(a, b string) int {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}

	for i := 0; i < minLen; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}

	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}

	return 0
}

func IsSorted(f func(a, b string) int, arr []string) bool {
	if len(arr) < 2 {
		return true
	}

	order := f(arr[0], arr[1])

	for i := 0; i < len(arr)-1; i++ {
		comparison := f(arr[i], arr[i+1])

		if (order > 0 && comparison < 0) || (order < 0 && comparison > 0) {
			return false
		}
	}

	return true
}
