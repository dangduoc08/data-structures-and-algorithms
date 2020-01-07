package prefixTree

func (childNode *PrefixTree) Contains(word string) bool {
	var lastIndex int = len(word) - 1
	var currentIndex int = 0
	// Default is false
	// if for loop not satisfied always return false
	var matched bool

	// Run loop for word length
	for childNode.Leaf[string([]rune(word)[currentIndex])] != nil {
		childNode = childNode.Leaf[string([]rune(word)[currentIndex])]
		// If not matched whole word
		// return false
		if currentIndex == lastIndex {
			matched = childNode.IsEndOfWord
			break
		}
		currentIndex++
	}
	return matched
}
