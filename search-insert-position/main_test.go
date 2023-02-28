package searchinsertposition

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums1 := []int{1, 3, 5, 6}
	target1 := 5
	expected1 := 2
	output1 := searchInsert(nums1, target1)
	if output1 != expected1 {
		t.Errorf("Test Case 1: Ouput %v - Expected %v", output1, expected1)
	}

	nums2 := []int{1, 3, 5, 6}
	target2 := 2
	expected2 := 1
	output2 := searchInsert(nums2, target2)
	if output2 != expected2 {
		t.Errorf("Test Case 2: Ouput %v - Expected %v", output2, expected2)
	}

	nums3 := []int{1, 3, 5, 6}
	target3 := 7
	expected3 := 4
	output3 := searchInsert(nums3, target3)
	if output3 != expected3 {
		t.Errorf("Test Case 3: Ouput %v - Expected %v", output3, expected3)
	}
}
