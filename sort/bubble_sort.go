package sort

type bubbleSortCallback func(a, b int) bool

func BubbleSort(slice *[]int, callback bubbleSortCallback) {
	var sliceLength int = len(*slice)
	for index, _ := range *slice {
		for i := 0; i < sliceLength-index; i++ {
			var j int = i + 1
			if j >= sliceLength {
				j = i
			}
			compareOperation := (*slice)[i] > (*slice)[j]
			if callback != nil {
				compareOperation = callback((*slice)[i], (*slice)[j])
			}
			if compareOperation {
				temp := (*slice)[i]
				(*slice)[i] = (*slice)[j]
				(*slice)[j] = temp
			}
		}
	}
}
