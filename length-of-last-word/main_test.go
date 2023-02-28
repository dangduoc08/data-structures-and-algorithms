package lengthoflastword

import (
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	case1 := "Hello World"
	expected1 := 5
	output1 := lengthOfLastWord(case1)
	if output1 != expected1 {
		t.Errorf("Test Case 1: Ouput %v - Expected %v", output1, expected1)
	}

	case2 := "   fly me   to   the moon  "
	expected2 := 4
	output2 := lengthOfLastWord(case2)
	if output2 != expected2 {
		t.Errorf("Test Case 2: Ouput %v - Expected %v", output2, expected2)
	}

	case3 := "luffy is still joyboy"
	expected3 := 6
	output3 := lengthOfLastWord(case3)
	if output3 != expected3 {
		t.Errorf("Test Case 3: Ouput %v - Expected %v", output3, expected3)
	}
}
