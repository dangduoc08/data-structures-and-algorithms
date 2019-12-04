package singlylinkedlist

func (note *Node) Length() int {
	currentNode := note
	var length int = 1
	for currentNode.Next != nil {
		currentNode = currentNode.Next
		length++
	}
	return length
}
