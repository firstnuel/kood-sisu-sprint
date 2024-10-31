package sprint

func AlphabetMastery(n int) string {
	acc := ""
	for i := 97; i < (97 + n); i++ {
		acc += string(i)
	}
	return acc
}
