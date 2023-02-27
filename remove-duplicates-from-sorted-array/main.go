package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	m := make(map[int]bool)
	res := 0
	i := 0
	for _, n := range nums {
		if !m[n] {
			m[n] = true
			res++
		} else {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			continue
		}
		i++
	}

	return res
}
