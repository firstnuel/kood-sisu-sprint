package sprint

import "fmt"

func Combinations() string {
	acc := ""
	for i := 0; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
				acc += fmt.Sprintf("%v%v%v", i, j, k)
				if !(i == 7 && j == 8 && k == 9) {
					acc += ", "
				}
			}
		}
	}
	return acc
}
