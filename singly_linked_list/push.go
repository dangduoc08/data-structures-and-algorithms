package singlylinkedlist

func (node *Node) Push(data interface{}) {
	if node.Data == nil {
		node.Data = data
		return
	}
	newNode := Node{
		Data: data,
	}
	for node.Next != nil {
		node = node.Next
		if node.Next == nil {
			node.Next = &newNode
			return
		}
	}
	if node.Next == nil {
		node.Next = &newNode
		return
	}
}
