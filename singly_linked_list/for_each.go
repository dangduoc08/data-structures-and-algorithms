package singlylinkedlist

type Func func(index int, currentNode *Node)

func (node *Node) ForEach(callback Func) {
	var index int = 0
	var length int = node.Length()
	for index < length {
		if callback != nil {
			callback(index, node)
		}
		node = node.Next
		index++
	}
}
