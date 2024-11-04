package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {

	if from > to {
		from, to = to, from
	}
	if to > len(arr)-1 {
		if from < 0 || from > len(arr) {
			s := make([]float64, 0)
			return s
		}
		return arr[:from]
	}
	if from < 0 || from > len(arr) {
		s := make([]float64, 0)
		return s
	}
	newArr := append(arr[:from], arr[to:]...)
	return newArr
}
