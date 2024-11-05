package sprint

func StrReverse(s string) string {
	rvrStr := ""
	for i := len(s) - 1; i >= 0; i-- {
		rvrStr += string(s[i])
	}
	return rvrStr
}
