package palindromenumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	case1 := 121
	expected1 := true
	output1 := isPalindrome(case1)
	if output1 != expected1 {
		t.Errorf("Test Case 1: Ouput %v - Expected %v", output1, expected1)
	}

	case2 := -121
	expected2 := false
	output2 := isPalindrome(case2)
	if output2 != expected2 {
		t.Errorf("Test Case 2: Ouput %v - Expected %v", output2, expected2)
	}

	case3 := 10
	expected3 := false
	output3 := isPalindrome(case3)
	if output3 != expected3 {
		t.Errorf("Test Case 3: Ouput %v - Expected %v", output3, expected3)
	}

	case4 := 1000021
	expected4 := false
	output4 := isPalindrome(case4)
	if output4 != expected4 {
		t.Errorf("Test Case 4: Ouput %v - Expected %v", output4, expected4)
	}

	case5 := 123
	expected5 := false
	output5 := isPalindrome(case5)
	if output5 != expected5 {
		t.Errorf("Test Case 5: Ouput %v - Expected %v", output5, expected5)
	}
}
