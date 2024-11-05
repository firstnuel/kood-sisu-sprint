package sprint

func SubstrIndex(s string, toFind string) int {
	if s == toFind {
		return 0
	}
	n := len(toFind)
	for i := 0; i < len(s)-n; i++ {
		if s[i:i+n] == toFind {
			return i
		}
	}
	return -1
}
