package singlylinkedlist

func (linkedList *Node) Splice(startIndex, deleteCount int, data interface{}) {
	var deleteChecker int = 1
	var splicedHeadNode *Node
	var splicedTailNode *Node

	linkedList.ForEach(func(currentNode *Node, index int) {
		if index >= startIndex-1 && startIndex != 0 {
			if deleteCount > 0 {
				if deleteChecker == 1 {
					splicedHeadNode = currentNode
				}
				if deleteChecker == deleteCount+2 {
					splicedTailNode = currentNode
				}
				deleteChecker++
			}
		} else if startIndex == 0 {
			if deleteCount > 0 && index == deleteCount {
				splicedHeadNode = currentNode
				splicedTailNode = currentNode.Next
			}
		}
	})
	splicedHeadNode.Next = splicedTailNode
}
