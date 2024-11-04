package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	if from > to {
		from, to = to, from
	}
	if from < 0 {
		from = 0
	}
	if to > len(arr) {
		to = len(arr)
	}
	if from >= len(arr) || to <= 0 {
		return []float64{}
	}

	return append(arr[:from], arr[to:]...)
}
