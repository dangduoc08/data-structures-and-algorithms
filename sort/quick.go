package sort

type quickCallback func(pivot, currentElem int) bool

func handlePartition(partition []int, sliceLength int, callback quickCallback) ([]int, []int) {
	var pivotIndex int = sliceLength - 1
	var pivot int = partition[pivotIndex]
	var j int = 0

	for i, currentElem := range partition {
		var compareOperation bool = pivot > currentElem
		if callback != nil {
			compareOperation = callback(pivot, currentElem)
		}
		if compareOperation {
			partition[i], partition[j] = swap(partition[i], partition[j])
			j++
		}
	}
	partition[j], partition[pivotIndex] = swap(partition[j], partition[pivotIndex])
	return partition[:j], partition[j+1:]
}

func Quick(slice []int, callback quickCallback) {
	var sliceLength int = len(slice)
	if sliceLength <= 1 {
		return
	}
	partition1, partition2 := handlePartition(slice, sliceLength, callback)
	Quick(partition1, callback)
	Quick(partition2, callback)
}

// 10 9 8 7 6 5 4 3 2 1

//
