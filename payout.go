package sprint

func Payout(amount int, denominations []int) (payout []int) {
	if amount <= 0 {
		return payout
	}
	denominations = SortInt(denominations)
	n := len(denominations) - 1
	for amount >= 0 {
		if amount-denominations[n] >= 0 {
			amount = amount - denominations[n]
			payout = append(payout, denominations[n])
		} else {
			n--
		}
		if n < 0 {
			break
		}
	}

	return payout
}

func SortInt(arr []int) []int {
	n := len(arr)
	for {
		swapped := false

		for i := 0; i < n-1; i++ {

			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		n--
	}
	return arr
}
