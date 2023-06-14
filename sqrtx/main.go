package sqrtx

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	if x == 1 {
		return 1
	}

	i := 1
	for i < x {
		if i*i < x {
			i++
		} else if i*i == x {
			return i
		} else {
			return i - 1
		}
	}
	return i - 1
}
