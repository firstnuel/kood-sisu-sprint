package sprint

func SortWordArr(a []string) []string {

	n := len(a)
	for {
		swapped := false
		for i := 0; i < n-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		n--
	}
	return a
}
