package sprint

func Overlap(arr1, arr2 []int) []int {
	// Create a map to store the counts of elements in arr1
	counts := make(map[int]int)
	for _, num := range arr1 {
		counts[num]++
	}

	// Iterate through arr2 and add common elements to the result
	var result []int
	for _, num := range arr2 {
		if counts[num] > 0 {
			result = append(result, num)
			counts[num]--
		}
	}

	// Sort the result using a custom sorting algorithm
	sortedResult := sortIntSlice(result)
	return sortedResult
}

func sortIntSlice(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}

	mid := len(slice) / 2
	left := sortIntSlice(slice[:mid])
	right := sortIntSlice(slice[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}
