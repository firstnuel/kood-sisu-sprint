package sprint

func sumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func FilterBySum(arr [][]int, limit int) [][]int {
	newArr := [][]int{}

	for _, arrItem := range arr {
		if sumArray(arrItem) >= limit {
			newArr = append(newArr, arrItem)
		}
	}
	return newArr
}
