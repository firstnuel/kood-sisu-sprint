package sprint

func StrConcatWith(strs []string, sep string) string {
	lastStr := strs[len(strs)-1]
	str := ""
	for _, v := range strs {
		if v != lastStr {
			str += v + sep
		}
	}
	return str + lastStr
}
