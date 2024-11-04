package sprint

func BulkAtoi(arr []string) []int {
	newArr := make([]int, len(arr))

	for i, v := range arr {
		newArr[i] = StrToInt(v)
	}

	return newArr
}
