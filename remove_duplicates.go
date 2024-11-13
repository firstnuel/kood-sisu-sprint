package sprint

func in(list []int, value int) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func RemoveDuplicates(arr []int) []int {
	newArr := []int{}

	for _, v := range arr {
		if !in(newArr, v) {
			newArr = append(newArr, v)
		}
	}
	return newArr
}
