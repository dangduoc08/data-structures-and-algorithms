package prefixTree

// Insert string to prefix tree using recursion, can use for loop also
func (childNode *PrefixTree) RecursiveInsert(word string) {
	var wordLength int = len(word)
	var isEndOfWord bool = wordLength == 0
	var runeStr []rune = []rune(word)
	var firstStr string
	var remainStr string

	// If not last word
	// to void panic error
	if !isEndOfWord {
		firstStr = string(runeStr[0])

		// Substring from index 1
		// remove first index
		// to pass to recursive function
		remainStr = string(runeStr[1:wordLength])
	}

	// Initialize child node for firsttime
	if childNode.Leaf == nil {
		childNode.Leaf = make(map[string]*PrefixTree)
	}

	if childNode.Leaf[firstStr] == nil && firstStr != "" {
		childNode.Leaf[firstStr] = &PrefixTree{}
	}

	if !childNode.IsEndOfWord {
		childNode.IsEndOfWord = isEndOfWord
	}

	// Condition to stop recursion
	if !isEndOfWord {
		childNode.Leaf[firstStr].RecursiveInsert(remainStr)
	}
}

func (childNode *PrefixTree) Insert(word string) {
	for _, runeStr := range word {
		var str string = string(runeStr)

		if childNode.Leaf == nil {
			childNode.Leaf = make(map[string]*PrefixTree)
		}

		if childNode.Leaf[str] == nil {
			childNode.Leaf[str] = new(PrefixTree)
		}

		if !childNode.IsEndOfWord {
			childNode.IsEndOfWord = false
		}
		childNode = childNode.Leaf[str]
	}
	childNode.IsEndOfWord = true
}
