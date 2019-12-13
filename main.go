package main

import (
	singlyLinkedList "github.com/dangduoc08/data-structures-and-algorithms/singly_linked_list"
	"github.com/dangduoc08/data-structures-and-algorithms/sort"
)

func main() {
	linkedList := singlyLinkedList.Node{}
	linkedList.Push(0)
	linkedList.Push(1)
	linkedList.Push(2)
	linkedList.Push(3)
	linkedList.Push(4)
	linkedList.Push(5)
	linkedList.Push(6)
	linkedList.Push(7)
	linkedList.Push(8)
	linkedList.Push(9)
	linkedList.Push(10)
	linkedList.Splice(0, 11, 11, 12, 13)

	slice := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort.Insertion(&slice, func(a, b int) bool {
		return a > b
	})

	slice1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort.Bubble(&slice1, func(a, b int) bool {
		return a > b
	})

	slice2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort.Selection(&slice2, func(a, b int) bool {
		return a > b
	})
}
