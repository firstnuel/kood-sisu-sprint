package sprint

func BalanceOut(arr []bool) []bool {
	trueCount, falseCount := 0, 0

	for _, v := range arr {
		if v {
			trueCount++
		} else {
			falseCount++
		}
	}

	for trueCount != falseCount {
		if trueCount < falseCount {
			arr = append(arr, true)
			trueCount++
		} else {
			arr = append(arr, false)
			falseCount++
		}
	}
	return arr
}
