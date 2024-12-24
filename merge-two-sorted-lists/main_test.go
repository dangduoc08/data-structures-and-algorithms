package mergetwosortedlists

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	// case1List1 := insert(new(ListNode), []int{1, 2, 4, 4})
	// case1List2 := insert(new(ListNode), []int{1, 3, 4})
	// expected1 := insert(new(ListNode), []int{1, 1, 2, 3, 4, 4})
	// output1 := mergeTwoLists(case1List1, case1List2)
	// traverse(output1, func(i int) {
	// 	fmt.Println(i)
	// })
	// if output1 != expected1 {
	// 	t.Errorf("Test Case 1: Ouput %v - Expected %v", output1, expected1)
	// }

	traverse(insert(new(ListNode), []int{0}), func(i int) {
		fmt.Println(i)
	})

	// case2List1 := new(ListNode)
	// case2List2 := new(ListNode)
	// expected2 := new(ListNode)
	// output2 := mergeTwoLists(case2List1, case2List2)
	// if output2 != expected2 {
	// 	t.Errorf("Test Case 2: Ouput %v - Expected %v", output2, expected2)
	// }

	// case3List1 := new(ListNode)
	// case3List2 := new(ListNode).insert([]int{0})
	// expected3 := new(ListNode)
	// output3 := mergeTwoLists(case3List1, case3List2)
	// if output3 != expected3 {
	// 	t.Errorf("Test Case 3: Ouput %v - Expected %v", output3, expected3)
	// }
}
