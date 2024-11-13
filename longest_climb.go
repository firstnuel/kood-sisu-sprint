package sprint

func In(list []int, value int) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func LongestClimb(arr []int) []int {
	newArr := []int{arr[0]}

	for _, v := range arr {
		if !In(newArr, v) {
			if v < newArr[0] {
				newArr = append([]int{v}, newArr...)
			} else {
				newArr = append(newArr, v)
			}
		}
	}
	return newArr
}
