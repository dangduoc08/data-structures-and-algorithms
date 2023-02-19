package findnumberswithevennumberofdigits

import "testing"

func TestFindNumbers(t *testing.T) {
	case1 := []int{12, 345, 2, 6, 7896}
	expected1 := 2
	output1 := findNumbers(case1)
	if output1 != expected1 {
		t.Errorf("Test Case 1: Ouput %v - Expected %v", output1, expected1)
	}

	case2 := []int{555, 901, 482, 1771}
	expected2 := 1
	output2 := findNumbers(case2)
	if output2 != expected2 {
		t.Errorf("Test Case 2: Ouput %v - Expected %v", output2, expected2)
	}
}
