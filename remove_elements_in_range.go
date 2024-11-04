package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {

	if from > to {
		from, to = to, from
	}
	if to > len(arr)-1 {
		return arr[:to]
	}
	newArr := append(arr[:from], arr[to:]...)
	return newArr
}
