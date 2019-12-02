package main

import (
	"fmt"
	singlyLinkedList "github.com/dangduoc08/data-structures-and-algorithms/singly_linked_list"
)

func main() {
	linkedList := singlyLinkedList.Create(1)
	singlyLinkedList.Append(&linkedList, 2)
	singlyLinkedList.Append(&linkedList, 3)
	singlyLinkedList.Append(&linkedList, 4)
	singlyLinkedList.Append(&linkedList, 5)
	singlyLinkedList.Append(&linkedList, 6)
	fmt.Println(linkedList)
}
