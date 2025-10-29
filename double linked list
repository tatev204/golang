package main

import (
	"fmt"
)

type DLL interface {
	Insert(val int)
	PrintListMoveForward()
	PrintListMoveBack()
	CountofNodes() int
	NegativeElementsbool() bool
	SumNodes() int
	SumOddIndexes() int
}
type Node struct {
	value int
	next  *Node
	prior *Node
}
type DoubleLinkedList struct {
	root *Node
	tail *Node
}

func (dll *DoubleLinkedList) Insert(val int) {
	node := &Node{value: val}
	if dll.root == nil {
		dll.root = node
		dll.tail = node
	} else {
		dll.tail.next = node
		node.prior = dll.tail
		dll.tail = node
	}
}

func (dll *DoubleLinkedList) PrintListMoveForward() {
	current := dll.root
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}
func (dll *DoubleLinkedList) PrintListMoveBack() {
	current := dll.tail
	for current != nil {
		fmt.Println(current.value)
		current = current.prior
	}

}
func (dll *DoubleLinkedList) CountofNodes() int {
	count := 0
	current := dll.root
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (dll *DoubleLinkedList) NegativeElementsbool() bool {
	current := dll.root

	for current != nil {
		if current.value < 0 {
			return true
		}

		current = current.next
	}
	return false
}

func (dll *DoubleLinkedList) SumNodes() int {
	sum := 0
	current := dll.root
	for current != nil {
		sum += current.value
		current = current.next
	}
	return sum
}
func (dll *DoubleLinkedList) SumOddIndexes() int {
	sumodd := 0
	count := dll.CountofNodes()
	current := dll.root

	for i := 1; i <= count; i += 2 {
		if current != nil {
			sumodd += current.value
		}
		current = current.next
	}
	return sumodd
}

func main() {
	var DoubleLinkedl DLL = &DoubleLinkedList{}

	DoubleLinkedl.Insert(1)
	DoubleLinkedl.Insert(2)

	fmt.Println("\nlist forward")
	DoubleLinkedl.PrintListMoveForward()
	fmt.Println("\nlist backward")
	DoubleLinkedl.PrintListMoveBack()

	fmt.Println("count of nodes՝", DoubleLinkedl.CountofNodes())
	fmt.Println("negativ element՝", DoubleLinkedl.NegativeElementsbool())
	fmt.Println("sum of nodes՝", DoubleLinkedl.SumNodes())
	fmt.Println("sum odd index՝", DoubleLinkedl.SumOddIndexes())

}
