package sprint

func StrLength(s string) []int {
	bytes := []byte(s)
	runes := []rune(s)

	return []int{len(runes), len(bytes)}
}
