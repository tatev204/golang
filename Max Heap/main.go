package main

import (
	"fmt"
)

type Node struct {
	value int
	right *Node
	left  *Node
}
type MaxHeap struct {
	root *Node
}

func (h *MaxHeap) insert(val int) {
	node := &Node{value: val}
	if h.root == nil {
		h.root = node
	}
	if val < h.root {
		h.root.right = node
	} else {
		h.root = node
	}
	if val < h.root {
		if val > h.root.right {
			h.root.left = node
		} else {
			h.root
		}

	}

}
