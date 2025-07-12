package main

import (
	"fmt"
)

type LL interface {
	insert(val int)
	CountofNodes() int
	NegativeElementsbool() bool
	SumNodes() int
	SumOddIndexes() int
	PrintList()
}

type Node struct {
	value int
	next  *Node
}
type LinkedList struct {
	root *Node
}

func (ll *LinkedList) insert(val int) {
	node := &Node{value: val}
	if ll.root == nil {
		ll.root = node
	} else {
		current := ll.root
		if current.next != nil {
			current = current.next
		}
		current.next = node
	}
}
func (ll *LinkedList) PrintList() {
	current := ll.root
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}
func (ll *LinkedList) CountofNodes() int {
	count := 0
	current := ll.root
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (ll *LinkedList) NegativeElementsbool() bool {
	current := ll.root

	for current != nil {
		if current.value < 0 {
			return true
		}

		current = current.next
	}
	return false
}

func (ll *LinkedList) SumNodes() int {
	sum := 0
	current := ll.root
	for current != nil {
		sum += current.value
		current = current.next
	}
	return sum
}
func (ll *LinkedList) SumOddIndexes() int {
	sumodd := 0
	count := ll.CountofNodes()
	current := ll.root

	for i := 1; i <= count; i += 2 {
		if current != nil {
			sumodd += current.value
		}
		current = current.next
	}
	return sumodd
}

func main() {
	var LinkedL LL = &LinkedList{}

	LinkedL.insert(1)
	LinkedL.insert(2)

	fmt.Println("\nlist ")
	LinkedL.PrintList()

	fmt.Println("count of nodes՝", LinkedL.CountofNodes())
	fmt.Println("negativ element՝", LinkedL.NegativeElementsbool())
	fmt.Println("sum of nodes՝", LinkedL.SumNodes())
	fmt.Println("sum odd index՝", LinkedL.SumOddIndexes())

}
