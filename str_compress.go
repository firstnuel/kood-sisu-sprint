package sprint

import "fmt"

func StrCompress(input string) string {
	if len(input) == 0 {
		return ""
	}

	count := 1
	firstCh := input[0]
	str := ""

	for i := 1; i < len(input); i++ {
		if input[i] == firstCh {
			count++
		} else {
			if count > 1 {
				str += fmt.Sprintf("%d", count) + string(firstCh)
			} else {
				str += string(firstCh)
			}
			firstCh = input[i]
			count = 1
		}
	}

	if count > 1 {
		str += fmt.Sprintf("%d", count) + string(firstCh)
	} else {
		str += string(firstCh)
	}
	return str
}
