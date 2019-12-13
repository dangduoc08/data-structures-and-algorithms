package sort

type selectionCallback func(a, b int) bool

func Selection(slice *[]int, callback selectionCallback) {
	var sliceValue []int = *slice
	var sliceLength int = len(sliceValue)
	for i, _ := range sliceValue {
		var flag int = sliceValue[i]
		var index int = i
		for j := i; j < sliceLength; j++ {
			var elem int = sliceValue[j]
			compareOperation := sliceValue[j] < flag
			if callback != nil {
				compareOperation = callback(flag, sliceValue[j])
			}
			if compareOperation {
				flag = elem
				index = j
			}
			if j == sliceLength-1 {
				var temp int = sliceValue[i]
				sliceValue[i] = flag
				sliceValue[index] = temp
			}
		}
	}
}
