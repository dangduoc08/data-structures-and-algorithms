package longestcommonprefix

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	case1 := []string{"flower", "flow", "flight"}
	expected1 := "fl"
	output1 := longestCommonPrefix(case1)
	if output1 != expected1 {
		t.Errorf("Test Case 1: Ouput %v - Expected %v", output1, expected1)
	}

	case2 := []string{"dog", "racecar", "car"}
	expected2 := ""
	output2 := longestCommonPrefix(case2)
	if output2 != expected2 {
		t.Errorf("Test Case 2: Ouput %v - Expected %v", output2, expected2)
	}
}
