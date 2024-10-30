package sprint

func BetweenLimits(from, to rune) string {
	acc := ""
	for i := to + 1; i < from; i++ {
		acc += string(i)
	}
	return acc
}
