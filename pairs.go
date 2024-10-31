package sprint

import "fmt"

func Pairs() string {
	acc := ""
	for i := 0; i < 99; i++ {
		for j := 1; j < 100; j++ {
			acc += fmt.Sprintf("%02d %02d", i, j)
			if !(i == 98 && j == 99) {
				acc += ", "
			}
		}
	}
	return acc
}
