package singlylinkedlist

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList interface {
	Append(d interface{})
}

func (n *Node) Append(d interface{}) {
	if n.Data == nil {
		n.Data = d
		return
	}

	newNode := Node{Data: d}

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
