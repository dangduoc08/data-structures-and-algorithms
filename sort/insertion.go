package sort

type insertionCallback func(prevElem, currentElem int) bool

func Insertion(slice []int, callback insertionCallback) {
	for i, _ := range slice {
		if i > 0 {
			var j int = i
			var prevElem int = slice[j-1]
			var currentElem int = slice[j]
			var compareOperation bool
			if callback != nil {
				compareOperation = callback(prevElem, currentElem)
			} else {
				compareOperation = prevElem > currentElem
			}
			for compareOperation {
				slice[j-1], slice[j] = swap(prevElem, currentElem)
				j--
				if j <= 0 {
					break
				}
				prevElem = slice[j-1]
				currentElem = slice[j]
				if callback != nil {
					compareOperation = callback(prevElem, currentElem)
				} else {
					compareOperation = prevElem > currentElem
				}
			}
		}
	}
}
