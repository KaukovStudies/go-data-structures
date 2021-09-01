package heap

type MaxHeap struct {
	items []int
}

func (h *MaxHeap) Insert(item int) {
	h.items = append(h.items, item)

	h.heapifyUp(len(h.items) - 1)
}

func (h *MaxHeap) Extract() int {
	extracted := h.items[0]

	itemLength := len(h.items) - 1

	if itemLength == 0 {
		// Heap is empty. Nothing to extract
		return -1
	}

	h.items[0] = h.items[itemLength]

	h.items = h.items[:itemLength]

	h.heapifyDown(0)

	return extracted
}

func (h *MaxHeap) heapifyUp(index int) {
	for h.items[parent(index)] < h.items[index] {
		h.swapItems(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) heapifyDown(index int) {
	lastIndex := len(h.items) - 1

	l, r := left(index), right(index)

	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex || h.items[l] > h.items[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.items[index] < h.items[childToCompare] {
			h.swapItems(index, childToCompare)

			index = childToCompare

			l, r = left(index), right(index)
		} else {
			// Early return since there's nothing more to do
			return
		}
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swapItems(i1, i2 int) {
	h.items[i1], h.items[i2] = h.items[i2], h.items[i1]
}
