package removeelement

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums1 := []int{3, 2, 2, 3}
	val1 := 3
	expected1 := 2
	output1 := removeElement(nums1, val1)
	if output1 != expected1 {
		t.Errorf("Test Case 1: Ouput %v - Expected %v", output1, expected1)
	}

	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val2 := 2
	expected2 := 5
	output2 := removeElement(nums2, val2)
	if output2 != expected2 {
		t.Errorf("Test Case 2: Ouput %v - Expected %v", output2, expected2)
	}
}
