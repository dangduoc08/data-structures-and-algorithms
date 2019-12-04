package singlylinkedlist

func (node *Node) Push(value interface{}) {
	if node.Value == nil {
		node.Value = value
		return
	}

	newNode := Node{Value: value}

	if node.Next == nil {
		node.Next = &newNode
		return
	}

	for node.Next != nil {
		node = node.Next
		if node.Next == nil {
			node.Next = &newNode
			return
		}
	}
}
