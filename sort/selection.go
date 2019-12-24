package sort

type selectionCallback func(flag, currentElem int) bool

func Selection(slice []int, callback selectionCallback) {
	var sliceLength int = len(slice)
	for i, _ := range slice {
		var flag int = slice[i] // flag = min || max
		for j := i; j < sliceLength; j++ {
			if j > 0 {
				var currentElem int = slice[j]
				var compareOperation bool
				if callback != nil {
					compareOperation = callback(flag, currentElem)
				} else {
					compareOperation = flag > currentElem
				}
				if compareOperation {
					slice[i], slice[j] = swap(flag, currentElem)
					flag = currentElem
				}
			}
		}
	}
}
