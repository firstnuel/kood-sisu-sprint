package sprint

func SortIntegerTable(table []int) []int {
	n := len(table)
	for {
		swapped := false
		for i := 0; i < n-1; i++ {
			if table[i] > table[i+1] {
				table[i], table[i+1] = table[i+1], table[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		n--
	}
	return table
}
