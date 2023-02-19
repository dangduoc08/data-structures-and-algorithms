package findnumberswithevennumberofdigits

func getIntLen(n int) int {
	res := 0
	for n > 0 {
		res += 1
		n = n / 10
	}
	return res
}

func findNumbers(nums []int) int {
	res := 0
	for _, num := range nums {
		if getIntLen(num)%2 == 0 {
			res += 1
		}
	}
	return res
}
