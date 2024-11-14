package sprint

func ToRoman(num int) string {
	roman := ""
	rf := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"D":  500,
		"M":  1000,
	}
	keys := []string{"I", "IV", "V", "X", "XL", "L", "XC", "C", "D", "M"}

	n := len(keys) - 1

	for num >= 0 {
		if num-rf[keys[n]] >= 0 {
			num = num - rf[keys[n]]
			roman += keys[n]
		} else {
			n--
		}
		if n < 0 {
			break
		}
	}
	return roman
}
