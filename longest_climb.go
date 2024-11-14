package sprint

func LongestClimb(arr []int) []int {
	start, count, lngstStart, lngstCount := 0, 0, 0, 0

	for i := 0; i < len(arr)-2; i++ {
		if arr[i] < arr[i+1] {
			count++
			start = i
			if count >= lngstCount {
				lngstStart = start - (count - 1)
			}
		} else if arr[i] > arr[i+1] {
			if count > lngstCount {
				lngstCount = count
			}
			count = 0
		}
	}
	if lngstStart+lngstCount >= len(arr)-2 {
		return arr[lngstStart:]
	}
	return arr[lngstStart : lngstStart+lngstCount+1]
}
