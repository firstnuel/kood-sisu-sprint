package sprint

func Payout(amount int, denominations []int) (payout []int) {
	if amount <= 0 {
		return payout
	}

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
