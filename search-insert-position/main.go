package searchinsertposition

func searchInsert(nums []int, target int) int {
	counter := 0
	if len(nums) == 1 {
		if target <= nums[0] {
			return 0
		} else {
			return 1
		}
	}

	for i, num := range nums {
		if num == target {
			return i
		} else if target > num {
			counter++
		} else {
			break
		}
	}

	return counter
}
