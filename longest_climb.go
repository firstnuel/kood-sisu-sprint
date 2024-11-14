package sprint

func LongestClimb(arr []int) []int {
	count, lngstCount := 0, 0
	lngstStart := 0

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			count++
			if count > lngstCount {
				lngstCount = count
				lngstStart = i - (count - 1)
			}
		} else {
			count = 0
		}
	}
	return arr[lngstStart : lngstStart+lngstCount+1]
}
