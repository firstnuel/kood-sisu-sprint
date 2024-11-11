package sprint

func NthFibonacci(index int) int {
	if index < 0 {
		return 0
	}
	if index == 1 {
		return 0
	}
	return 1 + NthFibonacci(index-1)
}
