package sprint

func BetweenLimits(from, to rune) string {
	if from > to {
		to, from = from, to
	}
	acc := ""
	for i := from + 1; i < to; i++ {
		acc += string(i)
	}
	return acc
}
