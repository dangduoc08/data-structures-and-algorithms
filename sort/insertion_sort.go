package sort

import "fmt"

func InsertionSort(slice *[]int) {
	var sliceLength int = len(*slice)
	for index, value := range *slice {
		var j int = index + 1
		if j >= sliceLength {
			j = index
		}
		for value > (*slice)[j] {
			fmt.Println(value)
			fmt.Println((*slice)[j])
			cloneValue := value
			(*slice)[index] = (*slice)[j]
			(*slice)[j] = cloneValue
		}
	}
}
