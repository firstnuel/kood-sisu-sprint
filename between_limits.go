package sprint

func BetweenLimits(from, to rune) string {
	acc := ""
	for i := from + 1; i < to; i++ {
		acc += string(i)
	}
	return acc
}
