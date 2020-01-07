package main

import (
	"fmt"

	prefixTree "github.com/dangduoc08/data-structures-and-algorithms/prefix_tree"
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
	sort.Insertion(slice, func(prev, current int) bool {
		return prev > current
	})

	slice1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort.Bubble(slice1, func(current, next int) bool {
		return current > next
	})

	slice2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort.Selection(slice2, func(flag, current int) bool {
		return flag > current
	})

	slice3 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort.Quick(slice3, func(pivot, current int) bool {
		return pivot < current
	})

	const (
		USERS      = "/users"
		PRODUCTS   = "/products"
		CATEGORIES = "/categories"
	)

	tree := prefixTree.PrefixTree{}
	tree.Insert(USERS)
	tree.Insert(PRODUCTS)
	var matched bool = tree.Contains(USERS)

	fmt.Println(matched)
}
