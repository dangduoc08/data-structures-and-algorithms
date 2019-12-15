package sort

type quickCallback func(a, b int) bool

func handlePartition(partition []int, sliceLength int) ([]int, []int) {
	var pivotIndex int = sliceLength - 1
	var pivot int = partition[pivotIndex]
	var j int = 0

	for i, elem := range partition {
		if elem < pivot {
			partition[i], partition[j] = swap(partition[i], partition[j])
			j++
		}
	}
	partition[j], partition[pivotIndex] = swap(partition[j], partition[pivotIndex])
	return partition[:j], partition[j+1:]
}

func Quick(slice []int) {
	var sliceLength int = len(slice)
	if sliceLength <= 1 {
		return
	}
	partition1, partition2 := handlePartition(slice, sliceLength)
	Quick(partition1)
	Quick(partition2)
}
