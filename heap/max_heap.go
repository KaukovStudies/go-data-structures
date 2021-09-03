package heap

import "errors"

type MaxHeap struct {
	items []int
}

func (h *MaxHeap) Insert(item int) {
	h.items = append(h.items, item)

	h.heapifyUp(len(h.items) - 1)
}

func (h *MaxHeap) Extract() (int, error) {
	itemLength := len(h.items)

	if len(h.items) == 0 {
		return -1, errors.New("heap is empty - nothing to extract")
	}

	extracted := h.items[0]

	h.items[0] = h.items[itemLength-1]

	h.items = h.items[:itemLength-1]

	h.heapifyDown(0)

	return extracted, nil
}

func (h *MaxHeap) swapItems(i1, i2 int) {
	h.items[i1], h.items[i2] = h.items[i2], h.items[i1]
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
