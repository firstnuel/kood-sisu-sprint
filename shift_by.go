package sprint

func ShiftBy(r rune, step int) rune {
	return 97 + (r-97+rune(step))%26
}
