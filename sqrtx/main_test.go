package sqrtx

import "testing"

func TestMySqrt(t *testing.T) {
	case1 := 4
	expected1 := 2
	output1 := mySqrt(case1)
	if output1 != expected1 {
		t.Errorf("Test Case 1: Ouput %v - Expected %v", output1, expected1)
	}

	case2 := 8
	expected2 := 2
	output2 := mySqrt(case2)
	if output2 != expected2 {
		t.Errorf("Test Case 2: Ouput %v - Expected %v", output2, expected2)
	}

	case3 := 2
	expected3 := 1
	output3 := mySqrt(case3)
	if output3 != expected3 {
		t.Errorf("Test Case 3: Ouput %v - Expected %v", output3, expected3)
	}

	case4 := 1
	expected4 := 1
	output4 := mySqrt(case4)
	if output4 != expected4 {
		t.Errorf("Test Case 4: Ouput %v - Expected %v", output4, expected4)
	}

	// case5 := 123
	// expected5 := 10
	// output5 := mySqrt(case5)
	// if output5 != expected5 {
	// 	t.Errorf("Test Case 5: Ouput %v - Expected %v", output5, expected5)
	// }
}
