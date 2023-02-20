package palindromenumber

func genReverseNum(x int) int {
	reverseX := 0
	for x > 0 {
		y := x % 10
		reverseX = (reverseX * 10) + y
		x = x / 10
	}
	return reverseX
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 9 {
		return true
	}

	return x == genReverseNum(x)
}
