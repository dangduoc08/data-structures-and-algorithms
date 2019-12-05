package main

import (
	"fmt"

	singlyLinkedList "github.com/dangduoc08/data-structures-and-algorithms/singly_linked_list"
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
	linkedList.ForEach(func(currentNode *singlyLinkedList.Node, index int) {
		fmt.Println(*currentNode)
	})
	fmt.Printf("Linked list length: %v \n", linkedList.Length())
}
