package twosum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	expected1 := []int{0, 1}
	output1 := twoSum(nums1, target1)
	if len(output1) == 0 || output1[0] != expected1[0] || output1[1] != expected1[1] {
		t.Errorf("Test Case 1: Ouput %v - Expected %v", output1, expected1)
	}

	nums2 := []int{3, 2, 4}
	target2 := 6
	expected2 := []int{1, 2}
	output2 := twoSum(nums2, target2)
	if len(output2) == 0 || output2[0] != expected2[0] || output2[1] != expected2[1] {
		t.Errorf("Test Case 2: Ouput %v - Expected %v", output2, expected2)
	}

	nums3 := []int{3, 3}
	target3 := 6
	expected3 := []int{0, 1}
	output3 := twoSum(nums3, target3)
	if len(output3) == 0 || output3[0] != expected3[0] || output3[1] != expected3[1] {
		t.Errorf("Test Case 3: Ouput %v - Expected %v", output3, expected3)
	}

	nums4 := []int{2, 5, 5, 11}
	target4 := 10
	expected4 := []int{1, 2}
	output4 := twoSum(nums4, target4)
	if len(output4) == 0 || output4[0] != expected4[0] || output4[1] != expected4[1] {
		t.Errorf("Test Case 4: Ouput %v - Expected %v", output4, expected4)
	}
}
