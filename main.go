package main

import (
	"fmt"
	singlyLinkedList "github.com/dangduoc08/data-structures-and-algorithms/singly_linked_list"
)

func main() {
	node := singlyLinkedList.Node{}
	node.Push(0)
	node.Push(1)
	node.Push(2)
	node.Push(3)
	node.Push(4)
	node.Push(5)
	node.Push(6)
	node.Push(7)
	node.Push(8)
	node.Push(9)
	node.Push(10)

	node.ForEach(func(index int, currentNode *singlyLinkedList.Node) {
		fmt.Printf("Index: %v has value: %v point to: %v \n", index, currentNode.Value, *currentNode)
	})
}
