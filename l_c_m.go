package sprint

func LCM(a, b int) int {
	multA, multB := []int{}, []int{}

	for i := 1; i <= 10; i++ {
		multA = append(multA, i*a)
		multB = append(multB, i*b)
	}

	for i := 0; i < len(multA); i++ {
		if Find(multA, multB[i]) > 0 {
			return multB[i]
		}
	}
	return 0
}

func Find(arr []int, target int) int {
	for _, v := range arr {
		if v == target {
			return target
		}
	}
	return 0
}
