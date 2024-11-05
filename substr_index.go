package sprint

func SubstrIndex(s string, toFind string) int {
	n := len(toFind)
	for i := 0; i < len(s)-n; i++ {
		if s[i:i+n] == toFind {
			return i
		}
	}
	return -1
}
