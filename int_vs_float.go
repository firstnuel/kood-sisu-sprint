package sprint

func IntVsFloat(i int, f float32) string {
	if i > int(f) {
		return "Integer"
	}
	return "Float"
}
