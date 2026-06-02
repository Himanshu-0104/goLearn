package main

type heap struct {
	data []int
	less func(a, b int) bool
}

// constructor for heap
func NewHeap(less func(a, b int) bool) *heap {
	return &heap{
		data: []int{},
		less: less,
	}
}

func (h *heap) Insert(value int) {
	h.data = append(h.data, value)
	h.heapifyUp(len(h.data) - 1) //what this does => refactor heap
}

func (h *heap) heapifyUp(index int) {
	for index > 0 {  //logn
		parent := (index - 1) / 2
		if h.less(h.data[parent], h.data[index]) { //Swap only when for max heap => parent < child (a<b) , for min heap => parent > child (a>b)
			h.data[index], h.data[parent] = h.data[parent], h.data[index]
			index = parent
		} else {
			break
		}
	}
}

func (h *heap) Delete() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	maxValue := h.data[0]
	h.data[0] = h.data[len(h.data)-1] //last ele gets first
	h.data = h.data[:len(h.data)-1] //free size
	h.heapifyDown(0) // why zero => bcz delte from top(highest) so refactor down from top
	return maxValue, true
}

func (h *heap) heapifyDown(index int) {
	lastIndex := len(h.data) - 1
	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		largest := index

		if leftChild <= lastIndex && h.less(h.data[largest], h.data[leftChild]) {
			largest = leftChild
		}
		if rightChild <= lastIndex && h.less(h.data[largest], h.data[rightChild]) {
			largest = rightChild
		}
		if largest == index {
			break
		}
		h.data[index], h.data[largest] = h.data[largest], h.data[index]
		index = largest // due to this log(n) , bcz traverse only below this
	}
}

func (h *heap) Peek() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	return h.data[0], true
}

func (h *heap) Size() int {
	return len(h.data)
}

func (h *heap) IsEmpty() bool {
	return len(h.data) == 0
}	

/* // Heap Implementation

mh := NewHeap(func(a, b int) bool {
	return a < b // for max heap
	// return a > b // for min heap
})

mh.Insert(10)
mh.Insert(20)
mh.Insert(5)

maxValue, _ := mh.Delete() // maxValue will be 20 for max heap, 5 for min heap

*/