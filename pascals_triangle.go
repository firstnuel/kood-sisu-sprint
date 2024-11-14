package sprint

func PascalsTriangle(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}

	triangle := make([][]int, n)

	for i := 0; i < n; i++ {
		// Create a new row with length i + 1
		row := make([]int, i+1)
		row[0], row[i] = 1, 1 // First and last elements are 1

		// Fill in the middle elements of the row
		for j := 1; j < i; j++ {
			row[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}

		triangle[i] = row
	}

	return triangle
}
