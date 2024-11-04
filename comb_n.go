package sprint

import "fmt"

func CombN(n int) []string {
	return backtrack("", 0, n)
}

func backtrack(current string, start, n int) []string {
	if len(current) == n {
		return []string{current}
	}

	var result []string
	for i := start; i <= 9; i++ {

		result = append(result, backtrack(current+fmt.Sprint(i), i+1, n)...)
	}
	return result
}
