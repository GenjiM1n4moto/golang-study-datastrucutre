package heapgo

import (
	"sort"
)

// 小根堆
type heapInterface interface {
	sort.Interface // Swap, Len, Less
	Push(x any)
	Pop() any
}

type IntHeap []int

// Implementation of heapInterface
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	length := h.Len()
	x := old[length-1]
	*h = old[0 : length-1]
	return x
}

// Pop returns minimal element of the heap
func Pop(h heapInterface) any {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

// Init
// The complexity is O(n) where n = h.Len().
func Init(h heapInterface) {
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

// Remove the ith element and reorder the heap and return the original ith element
func Remove(h heapInterface, i int) any {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}

// Fix re-establish the heap ordering after the element of index i changes its value
func Fix(h heapInterface, i int) {
	if !down(h, i, h.Len()) { // if the down return true, it means the index's child's order has been changed, so the parent's order should be updated
		up(h, i)
	}
}

func up(h heapInterface, j int) {
	for {
		i := (j - 1) / 2 // find parent
		if i == 0 || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

// down makes the i0's children's order right
func down(h heapInterface, i0 int, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // limitation
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			// Find the smaller child for the following exchange to the parent
			j = j2
		}
		if !h.Less(j, i) { // if the parent is smaller, should not be changed
			break
		}
		h.Swap(i, j) // Exchange the position of parent and child
		i = j        // Continue exchanging
	}
	return i > i0
}
