package sprint

func Overlap(arr1, arr2 []int) []int {
	// Create a map to store the counts of elements in the first array
	counts := make(map[int]int)

	// Count the occurrences of elements in the first array
	for _, num := range arr1 {
		counts[num]++
	}

	// Create a slice to store the common elements
	var result []int

	// Iterate through the second array and add common elements to the result
	for _, num := range arr2 {
		if _, ok := counts[num]; ok {
			result = append(result, num)
			counts[num]--
			if counts[num] == 0 {
				delete(counts, num)
			}
		}
	}

	return result
}
