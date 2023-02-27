package removeelement

func removeElement(nums []int, val int) int {
	i := 0
	j := i

	for i < len(nums) {
		num := nums[i]

		if num != val {
			nums[j] = num
			j++
		}

		i++
	}

	return j
}
