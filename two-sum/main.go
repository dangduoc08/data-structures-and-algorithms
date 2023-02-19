package twosum

func twoSum(nums []int, target int) []int {
	numMap := map[int]int{}

	for i, num := range nums {
		if _, ok := numMap[target-num]; ok {
			return []int{numMap[target-num], i}
		} else {
			numMap[num] = i
		}
	}

	return []int{}
}
