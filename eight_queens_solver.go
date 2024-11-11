package sprint

import (
	"strconv"
	"strings"
)

// Helper function to check if a queen can be placed at (row, col) without conflicts
func isValid(queens []int, row, col int) bool {
	for c, r := range queens {
		// Check same row or diagonal
		if r == row || abs(r-row) == abs(c-col) {
			return false
		}
	}
	return true
}

// Recursive function to solve the eight queens puzzle
func solve(col int, queens []int, results *[]string) {
	if col == 8 { // All queens placed
		var solution strings.Builder
		for _, r := range queens {
			solution.WriteString(strconv.Itoa(r + 1))
		}
		*results = append(*results, solution.String())
		return
	}

	for row := 0; row < 8; row++ {
		if isValid(queens, row, col) {
			solve(col+1, append(queens, row), results)
		}
	}
}

// Absolute value function for integer difference
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Main function to return all solutions as a formatted string
func EightQueensSolver() string {
	var results []string
	solve(0, []int{}, &results)
	return strings.Join(results, "\n")
}
