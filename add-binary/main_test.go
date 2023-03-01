package addbinary

import (
	"testing"
)

func TestAddBinary(t *testing.T) {
	a1 := "11"
	b1 := "1"
	expected1 := "100"
	output1 := addBinary(a1, b1)
	if output1 != expected1 {
		t.Errorf("Test Case 1: Ouput %v - Expected %v", output1, expected1)
	}

	a2 := "1010"
	b2 := "1011"
	expected2 := "10101"
	output2 := addBinary(a2, b2)
	if output2 != expected2 {
		t.Errorf("Test Case 2: Ouput %v - Expected %v", output2, expected2)
	}

	a3 := "1111"
	b3 := "1111"
	expected3 := "11110"
	output3 := addBinary(a3, b3)
	if output3 != expected3 {
		t.Errorf("Test Case 3: Ouput %v - Expected %v", output3, expected3)
	}
}
