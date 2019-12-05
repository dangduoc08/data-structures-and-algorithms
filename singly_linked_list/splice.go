package singlylinkedlist

func createNewNode(data ...interface{}) *Node {
	var node *Node = &Node{}
	for _, value := range data {
		node.Push(value)
	}
	return node
}

func deleteNodeAtHead(node *Node, deleteCount int) (*Node, *Node) {
	var splicedHeadNode *Node
	var splicedTailNode *Node
	node.ForEach(func(currentNode *Node, index int) {
		if index == 0 {
			splicedHeadNode = currentNode
		}
		if deleteCount > 0 && index == deleteCount {
			splicedHeadNode.Data = currentNode.Data
			splicedTailNode = currentNode.Next
		}
	})
	return splicedHeadNode, splicedTailNode
}

func deleteNodeAtIndex(node *Node, startIndex, deleteCount int) (*Node, *Node) {
	var splicedHeadNode *Node
	var splicedTailNode *Node
	var deleteChecker int = 0
	node.ForEach(func(currentNode *Node, index int) {
		if index == startIndex-1 {
			splicedHeadNode = currentNode
		} else if deleteChecker-1 == deleteCount {
			splicedTailNode = currentNode
		}
		if index >= startIndex-1 {
			deleteChecker++
		}
	})
	return splicedHeadNode, splicedTailNode
}

func (node *Node) Splice(startIndex, deleteCount int, data ...interface{}) {
	var splicedHeadNode *Node = node
	var splicedTailNode *Node = node.Next
	var newNode *Node = nil
	// Create new node from data
	if data[0] != nil {
		newNode = createNewNode(data...)
	}

	// Delete node at index = 0 or specified index
	if deleteCount > 0 {
		if startIndex == 0 {
			splicedHeadNode, splicedTailNode = deleteNodeAtHead(node, deleteCount)
		} else if startIndex > 0 {
			splicedHeadNode, splicedTailNode = deleteNodeAtIndex(node, startIndex, deleteCount)
		}
		// Merge 2 linked lists to new one
		splicedHeadNode.Next = splicedTailNode
	}

	// Add node at index = 0 or specified index or at index = linked list length
	if newNode != nil {
		if startIndex == 0 {
			var lastTailIndex int = len(data) - 2
			temp := *splicedHeadNode
			splicedHeadNode.Data = newNode.Data
			splicedTailNode = newNode.Next
			splicedTailNode.ForEach(func(currentNode *Node, index int) {
				if index == lastTailIndex {
					currentNode.Next = &temp
				}
			})
		} else if startIndex <= node.Length()-1 {
			node.ForEach(func(currentNode *Node, index int) {
				if index == startIndex-1 {
					splicedHeadNode = currentNode
				}
			})
			splicedTailNode = newNode
			splicedTailNode.ForEach(func(currentNode *Node, index int) {
				if currentNode.Next == nil {
					currentNode.Next = splicedHeadNode.Next
				}
			})
		} else if startIndex >= node.Length() {
			node.ForEach(func(currentNode *Node, index int) {
				if currentNode.Next == nil {
					splicedHeadNode = currentNode
					splicedTailNode = newNode
				}
			})
		}
		// Merge 2 linked lists to new one
		splicedHeadNode.Next = splicedTailNode
	}
}
