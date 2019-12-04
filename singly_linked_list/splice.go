package singlylinkedlist

// import "fmt"

// func (node *Node) Splice(startIndex, deleteCount int, data interface{}) {
// 	var index int = 0
// 	var deleteChecker int = 1
// 	var splicedNode *Node
// 	prevNode := node
// 	curNode := node
// 	nxtNode := node

// 	for curNode.Next != nil {
// 		if index >= startIndex {
// 			if deleteCount > 0 && deleteChecker <= deleteCount {
// 				if index == startIndex && startIndex > 0 {
// 					splicedNode = prevNode
// 				}
// 				if index-1 == startIndex && startIndex == 0 {
// 					splicedNode = nxtNode
// 					fmt.Println(splicedNode)
// 				}
// 				if deleteChecker == deleteCount {
// 					splicedNode.Next = nxtNode
// 					fmt.Println(nxtNode)
// 				}
// 				deleteChecker++
// 			}
// 			if data != nil {

// 			}
// 		}
// 		prevNode = curNode
// 		curNode = curNode.Next
// 		nxtNode = curNode.Next
// 		index++
// 	}
// }
