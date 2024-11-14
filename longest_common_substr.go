package sprint

func LongestCommonSubstr(str1, str2 string) string {
	m, n := len(str1), len(str2)
	// Create a 2D slice to store lengths of longest common suffixes
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Variable to store the length and the end position of the longest common substring
	maxLength, endPos := 0, 0

	// Fill the dp array
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
					endPos = i
				}
			}
		}
	}

	// If no common substring, return an empty string
	if maxLength == 0 {
		return ""
	}

	// Extract the longest common substring
	return str1[endPos-maxLength : endPos]
}
