package removeduplicatesfromsortedarray

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	case1 := []int{1, 1, 2}
	expected1 := 2
	output1 := removeDuplicates(case1)
	if output1 != expected1 {
		t.Errorf("Test Case 1: Ouput %v - Expected %v", output1, expected1)
	}

	case2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expected2 := 5
	output2 := removeDuplicates(case2)
	if output2 != expected2 {
		t.Errorf("Test Case 2: Ouput %v - Expected %v", output2, expected2)
	}
}
