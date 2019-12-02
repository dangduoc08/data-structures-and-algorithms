package singlylinkedlist

type Node struct {
	Data interface{}
	Next *Node
}

func Create(d interface{}) Node {
	head := Node{
		Data: d,
		Next: nil,
	}
	return head
}
