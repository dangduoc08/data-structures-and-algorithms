package singlylinkedlist

type forEachCallback func(currentNode *Node, index int)

func (node *Node) ForEach(callback forEachCallback) {
	var index int = 0
	for node.Next != nil {
		callback(node, index)
		node = node.Next
		index++
	}
	if node.Next == nil {
		callback(node, index)
		return
	}
}
