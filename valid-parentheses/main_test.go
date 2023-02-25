package validparentheses

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	case1 := "()"
	expected1 := true
	output1 := isValid(case1)
	if output1 != expected1 {
		t.Errorf("Test Case 1: Ouput %v - Expected %v", output1, expected1)
	}

	case2 := "()[]{}"
	expected2 := true
	output2 := isValid(case2)
	if output2 != expected2 {
		t.Errorf("Test Case 2: Ouput %v - Expected %v", output2, expected2)
	}

	case3 := "(]"
	expected3 := false
	output3 := isValid(case3)
	if output3 != expected3 {
		t.Errorf("Test Case 3: Ouput %v - Expected %v", output3, expected3)
	}

	case4 := "(){}}{"
	expected4 := false
	output4 := isValid(case4)
	if output4 != expected4 {
		t.Errorf("Test Case 4: Ouput %v - Expected %v", output4, expected4)
	}

	case5 := "{()}{}[{()}]()"
	expected5 := true
	output5 := isValid(case5)
	if output5 != expected5 {
		t.Errorf("Test Case 5: Ouput %v - Expected %v", output5, expected5)
	}

	case6 := "({[)"
	expected6 := false
	output6 := isValid(case6)
	if output6 != expected6 {
		t.Errorf("Test Case 6: Ouput %v - Expected %v", output6, expected6)
	}

	// case7 := "(([]){})"
	case7 := "(([])){}"
	expected7 := true
	output7 := isValid(case7)
	if output7 != expected7 {
		t.Errorf("Test Case 7: Ouput %v - Expected %v", output7, expected7)
	}
}
