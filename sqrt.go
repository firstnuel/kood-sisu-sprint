package sprint

func Sqrt(n int) int {
	// Return 0 for numbers less than 1, as they don't have meaningful square roots in this context
	if n < 1 {
		return 0
	}

	// Iterate from 1 to n to find an integer square root
	for i := 1; i <= n/2+1; i++ {
		if i*i == n { // Check if i * i equals n
			return i // If it does, i is the integer square root
		} else if i*i > n { // If i * i exceeds n, no integer square root exists
			return 0
		}
	}
	return 0 // In case no integer square root is found
}
