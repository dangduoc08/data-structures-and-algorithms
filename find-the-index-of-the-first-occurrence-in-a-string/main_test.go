package findtheindexofthefirstoccurrenceinastring

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	input1_a := "sadbutsad"
	input1_b := "sad"
	expected1 := 0
	output1 := strStr(input1_a, input1_b)
	if output1 != expected1 {
		t.Errorf("Test Case 1: Ouput %v - Expected %v", output1, expected1)
	}

	input2_a := "leetcode"
	input2_b := "leeto"
	expected2 := -1
	output2 := strStr(input2_a, input2_b)
	if output2 != expected2 {
		t.Errorf("Test Case 2: Ouput %v - Expected %v", output2, expected2)
	}

	input3_a := "loislovesuperman"
	input3_b := "love"
	expected3 := 4
	output3 := strStr(input3_a, input3_b)
	if output3 != expected3 {
		t.Errorf("Test Case 3: Ouput %v - Expected %v", output3, expected3)
	}

	input4_a := "mississippi"
	input4_b := "issip"
	expected4 := 4
	output4 := strStr(input4_a, input4_b)
	if output4 != expected4 {
		t.Errorf("Test Case 4: Ouput %v - Expected %v", output4, expected4)
	}

	input5_a := "mississippi"
	input5_b := "issipi"
	expected5 := -1
	output5 := strStr(input5_a, input5_b)
	if output5 != expected5 {
		t.Errorf("Test Case 5: Ouput %v - Expected %v", output5, expected5)
	}

	input6_a := "mississippi"
	input6_b := "pi"
	expected6 := 9
	output6 := strStr(input6_a, input6_b)
	if output6 != expected6 {
		t.Errorf("Test Case 6: Ouput %v - Expected %v", output6, expected6)
	}

	input7_a := "abbc"
	input7_b := "bbb"
	expected7 := -1
	output7 := strStr(input7_a, input7_b)
	if output7 != expected7 {
		t.Errorf("Test Case 7: Ouput %v - Expected %v", output7, expected7)
	}
}
