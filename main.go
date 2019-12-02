package main

import (
	"fmt"
	singlyLinkedList "github.com/dangduoc08/data-structures-and-algorithms/singly_linked_list"
)

func main() {
	node := singlyLinkedList.Node{}
	var linkedList singlyLinkedList.LinkedList
	linkedList = &node
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	fmt.Println(node.Next)
}
