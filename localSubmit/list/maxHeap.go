// package main

// type maxHeap struct {
// 	data []int
// }

// func (h *maxHeap) Insert(value int) {
// 	h.data = append(h.data, value)
// 	h.heapifyUp(len(h.data) - 1) //what this does => refactor heap
// }

// func (h *maxHeap) heapifyUp(index int) {
// 	for index > 0 {  //logn
// 		parent := (index - 1) / 2
// 		if h.data[index] > h.data[parent] {
// 			h.data[index], h.data[parent] = h.data[parent], h.data[index]
// 			index = parent
// 		} else {
// 			break
// 		}
// 	}
// }

// func (h *maxHeap) Delete() (int, bool) {
// 	if len(h.data) == 0 {
// 		return 0, false
// 	}
// 	maxValue := h.data[0]
// 	h.data[0] = h.data[len(h.data)-1] //last ele gets first
// 	h.data = h.data[:len(h.data)-1] //free size
// 	h.heapifyDown(0) // why zero => bcz delte from top(highest) so refactor down from top
// 	return maxValue, true
// }

// func (h *maxHeap) heapifyDown(index int) {
// 	lastIndex := len(h.data) - 1
// 	for {
// 		leftChild := 2*index + 1
// 		rightChild := 2*index + 2
// 		largest := index

// 		if leftChild <= lastIndex && h.data[leftChild] > h.data[largest] {
// 			largest = leftChild
// 		}
// 		if rightChild <= lastIndex && h.data[rightChild] > h.data[largest] {
// 			largest = rightChild
// 		}
// 		if largest == index {
// 			break
// 		}
// 		h.data[index], h.data[largest] = h.data[largest], h.data[index]
// 		index = largest // due to this log(n) , bcz traverse only below this
// 	}
// }

// func (h *maxHeap) Peek() (int, bool) {
// 	if len(h.data) == 0 {
// 		return 0, false
// 	}
// 	return h.data[0], true
// }

// func (h *maxHeap) Size() int {
// 	return len(h.data)
// }

// func (h *maxHeap) IsEmpty() bool {
// 	return len(h.data) == 0
// }	