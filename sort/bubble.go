package sort

type bubbleCallback func(a, b int) bool

func Bubble(slice *[]int, callback bubbleCallback) {
	var sliceValue []int = *slice
	var sliceLength int = len(sliceValue)
	for index, _ := range sliceValue {
		for i := 0; i < sliceLength-index; i++ {
			var j int = i + 1
			if j >= sliceLength {
				j = i
			}
			compareOperation := sliceValue[i] > sliceValue[j]
			if callback != nil {
				compareOperation = callback(sliceValue[i], sliceValue[j])
			}
			if compareOperation {
				temp := sliceValue[i]
				sliceValue[i] = sliceValue[j]
				sliceValue[j] = temp
			}
		}
	}
}
