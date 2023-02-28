package plusone

func plusOne(digits []int) []int {
	if len(digits) == 1 && digits[0] == 9 {
		return []int{1, 0}
	}

	i := 1
	lastDigit := digits[len(digits)-i]
	if lastDigit != 9 {
		digits[len(digits)-i] = lastDigit + 1
	} else {

		// last digit = 9
		for {
			digits[len(digits)-i] = 0
			if i == len(digits) {
				digits = append([]int{1}, digits...)
				break
			} else if digits[len(digits)-i-1] != 9 {
				digits[len(digits)-i-1] = digits[len(digits)-i-1] + 1
				break
			}
			i++
		}
	}

	return digits
}
