package sprint

func FromRoman(s string) int {

	sum := 0
	rf := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := 0; i < len(s)-1; i++ {
		if rf[string(s[i])] < rf[string(s[i+1])] {
			sum -= rf[string(s[i])]
		} else {
			sum += rf[string(s[i])]
		}
	}
	return sum + rf[string(s[len(s)-1])]
}
