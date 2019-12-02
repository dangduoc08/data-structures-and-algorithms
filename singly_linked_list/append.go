package singlylinkedlist

func Append(n *Node, d interface{}) {
	newNode := Node{
		Data: d,
		Next: nil,
	}

	if n.Next == nil {
		n.Next = &newNode
		return
	}

	for n.Next != nil {
		n = n.Next
		if n.Next == nil {
			n.Next = &newNode
			return
		}
	}
}
