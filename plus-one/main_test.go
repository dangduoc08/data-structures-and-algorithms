package plusone

import (
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	case1 := []int{1, 2, 3}
	expected1 := []int{1, 2, 4}
	output1 := plusOne(case1)
	if output1[0] != expected1[0] &&
		output1[1] != expected1[1] &&
		output1[2] != expected1[2] {
		t.Errorf("Test Case 1: Ouput %v - Expected %v", output1, expected1)
	}

	case2 := []int{4, 3, 2, 1}
	expected2 := []int{4, 3, 2, 2}
	output2 := plusOne(case2)
	if output2[0] != expected2[0] &&
		output2[1] != expected2[1] &&
		output2[2] != expected2[2] &&
		output2[3] != expected2[3] {
		t.Errorf("Test Case 2: Ouput %v - Expected %v", output2, expected2)
	}

	case3 := []int{9}
	expected3 := []int{1, 0}
	output3 := plusOne(case3)
	if output3[0] != expected3[0] &&
		output3[1] != expected3[1] {
		t.Errorf("Test Case 3: Ouput %v - Expected %v", output3, expected3)
	}

	case4 := []int{9, 9}
	expected4 := []int{1, 0, 0}
	output4 := plusOne(case4)
	if output4[0] != expected4[0] &&
		output4[1] != expected4[1] &&
		output4[2] != expected4[2] {
		t.Errorf("Test Case 4: Ouput %v - Expected %v", output4, expected4)
	}
}
