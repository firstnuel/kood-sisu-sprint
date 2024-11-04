package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {

	if from > to {
		from, to = to, from
	}
	if from < 0 || from > len(arr) {
		return nil
	}
	if to > len(arr)-1 {
		if from < 0 || from > len(arr) {
			return nil
		}
		return arr[:from]
	}
	newArr := append(arr[:from], arr[to:]...)
	return newArr
}
