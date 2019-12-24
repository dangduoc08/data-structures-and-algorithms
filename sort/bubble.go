package sort

type bubbleCallback func(currentElem, nextElem int) bool

func Bubble(slice []int, callback bubbleCallback) {
	var sliceLength int = len(slice)
	for index, _ := range slice {
		for i := 0; i < sliceLength-index; i++ {
			if i < sliceLength-1 {
				var j int = i + 1
				var currentElem int = slice[i]
				var nextElem int = slice[j]
				compareOperation := currentElem > nextElem
				if callback != nil {
					compareOperation = callback(currentElem, nextElem)
				}
				if compareOperation {
					slice[i], slice[j] = swap(currentElem, nextElem)
				}
			}
		}
	}
}
