package sprint

import "fmt"

func StrConcat(s1, s2, delim string) string {
	concatedStr := fmt.Sprintf("%v%v%v", s1, delim, s2)
	return concatedStr
}
