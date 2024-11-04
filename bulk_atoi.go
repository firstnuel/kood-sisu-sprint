package sprint

func strToInt(s string) int {
	start, acc := 0, 0
	IsNegative := false

	if s[0] == '-' {
		start = 1
		IsNegative = true
	}

	if s[0] == '+' {
		start = 1
	}

	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		if acc != 0 || s[i] != '0' {
			digit := int(s[i] - '0')
			acc = acc*10 + digit
		}
	}

	if IsNegative {
		return -acc
	}
	return acc
}

func BulkAtoi(arr []string) []int {
	newArr := make([]int, len(arr))

	for i, v := range arr {
		newArr[i] = strToInt(v)
	}

	return newArr
}
