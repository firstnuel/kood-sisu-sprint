package sprint

func AlphabetMastery(n int) string {
	acc := ""
	for i := 97; i <= (i + n); i++ {
		acc += string(i)
	}
	return acc
}
