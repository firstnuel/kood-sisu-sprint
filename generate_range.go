package sprint

func GenerateRange(min, max int) []int {

	if min > max {
		return nil
	}
	rangeSlice := make([]int, 0)

	idx := 0
	for i := min; i < max; i++ {
		rangeSlice = append(rangeSlice, i)
		idx++
	}
	return rangeSlice
}
