package palindromenumber

func forInt(x int, cb func(n int)) {
	for x > 0 {
		y := x % 10
		x = x / 10
		cb(y)
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 9 {
		return true
	}

	numArr := []int{}
	forInt(x, func(n int) {
		numArr = append(numArr, n)
	})
	l := len(numArr)

	for i, n := range numArr {
		if i <= (l/2)-1 {
			if n != numArr[l-i-1] {
				return false
			}
		} else {
			break
		}
	}

	return true
}
