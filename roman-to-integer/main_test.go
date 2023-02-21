package romantointeger

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	case1 := "III"
	expected1 := 3
	output1 := romanToInt(case1)
	if output1 != expected1 {
		t.Errorf("Test Case 1: Ouput %v - Expected %v", output1, expected1)
	}

	case2 := "LVIII"
	expected2 := 58
	output2 := romanToInt(case2)
	if output2 != expected2 {
		t.Errorf("Test Case 2: Ouput %v - Expected %v", output2, expected2)
	}

	case3 := "MCMXCIV"
	expected3 := 1994
	output3 := romanToInt(case3)
	if output3 != expected3 {
		t.Errorf("Test Case 3: Ouput %v - Expected %v", output3, expected3)
	}
}
