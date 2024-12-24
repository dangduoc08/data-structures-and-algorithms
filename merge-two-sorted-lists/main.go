package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func insert(l *ListNode, vals []int) *ListNode {
	head := l
	for i, val := range vals {
		l.Val = val
		if i < len(vals)-1 {
			l.Next = new(ListNode)
			l = l.Next
		}
	}

	return head
}

func traverse(l *ListNode, cb func(int)) *ListNode {
	head := l
	for l != nil {
		cb(l.Val)
		l = l.Next
	}

	return head
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	sortedArr := []int{}
	if list1 == nil && list2 == nil {
		return new(ListNode)
	}
	if list1 == nil && list2 != nil {
		return list2
	}
	if list1 != nil && list2 == nil {
		return list1
	}

	for list1 != nil || list2 != nil {
		if list1 == nil {
			traverse(list2, func(i int) {
				sortedArr = append(sortedArr, i)
			})
			break
		} else if list2 == nil {
			traverse(list1, func(i int) {
				sortedArr = append(sortedArr, i)
			})
			break
		} else if list1.Val < list2.Val {
			sortedArr = append(sortedArr, list1.Val)
			list1 = list1.Next
		} else if list1.Val > list2.Val {
			sortedArr = append(sortedArr, list2.Val)
			list2 = list2.Next
		} else if list1.Val == list2.Val {
			sortedArr = append(sortedArr, list1.Val, list2.Val)
			list1 = list1.Next
			list2 = list2.Next
		}
	}

	return insert(new(ListNode), sortedArr)
}
