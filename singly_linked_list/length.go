package singlylinkedlist

func (node *Node) Length() int {
	currentNode := node
	var length int = 1
	for currentNode.Next != nil {
		currentNode = currentNode.Next
		length++
	}
	return length
}
