package sprint

import "fmt"

func BetweenLimits(from, to rune) string {
	acc := ""
	for i := to + 1; i < from; i++ {
		str := fmt.Sprintf("%c", i)
		acc += str
	}
	return acc
}
