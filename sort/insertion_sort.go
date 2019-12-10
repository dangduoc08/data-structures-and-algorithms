package sort

type insertionSortCallback func(a, b int) bool

func InsertionSort(slice *[]int, callback insertionSortCallback) {
	var sliceValue []int = *slice
	for i, _ := range sliceValue {
		var j int = i
		if j > 0 {
			compareOperation := sliceValue[j] < sliceValue[j-1]
			if callback != nil {
				compareOperation = callback(sliceValue[j-1], sliceValue[j])
			}
			for j > 0 && compareOperation {
				var temp int = sliceValue[j-1]
				sliceValue[j-1] = sliceValue[j]
				sliceValue[j] = temp
				j--
			}
		}
	}
}
