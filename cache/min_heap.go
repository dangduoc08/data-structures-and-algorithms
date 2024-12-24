package cache

type MinHeap struct {
	list []int
}

func NewMinHeap(rootValue int) *MinHeap {
	h := &MinHeap{}
	h.list = append(h.list, rootValue)

	return h
}

func (h *MinHeap) getParentIndex(i int) int {
	if i == 0 {
		return 0
	}

	return (i - 1) / 2
}

func (h *MinHeap) getLeftIndex(i int) int {
	return i*2 + 1
}

func (h *MinHeap) getRightIndex(i int) int {
	return i*2 + 2
}

func (h *MinHeap) heapify() *MinHeap {
	i := 0
	if len(h.list) == 1 {
		return h
	}

	leftIndex := h.getLeftIndex(i)
	rightIndex := h.getRightIndex(i)
	swappableIndex := leftIndex
	if rightIndex < len(h.list) && h.list[leftIndex] > h.list[rightIndex] {
		swappableIndex = rightIndex
	}
	if swappableIndex >= len(h.list) {
		return h
	}

	for h.list[swappableIndex] <= h.list[i] {
		h.list[swappableIndex], h.list[i] = h.list[i], h.list[swappableIndex]

		i = swappableIndex
		leftIndex = h.getLeftIndex(i)
		rightIndex = h.getRightIndex(i)
		swappableIndex = leftIndex
		if rightIndex < len(h.list) && h.list[leftIndex] > h.list[rightIndex] {
			swappableIndex = rightIndex
		}
		if swappableIndex >= len(h.list) {
			return h
		}
	}

	return h
}

func (h *MinHeap) Insert(newValue int) *MinHeap {
	newIndex := len(h.list)
	parentIndex := h.getParentIndex(newIndex)
	h.list = append(h.list, newValue)

	for h.list[newIndex] < h.list[parentIndex] {
		h.list[newIndex], h.list[parentIndex] = h.list[parentIndex], newValue
		newIndex, parentIndex = parentIndex, h.getParentIndex(parentIndex)
	}

	return h
}

func (h *MinHeap) RemoveMin() *MinHeap {
	length := len(h.list)

	if length > 0 {
		h.list[0] = h.list[length-1]
		h.list = h.list[:length-1]
		h.heapify()
	}

	return h
}
