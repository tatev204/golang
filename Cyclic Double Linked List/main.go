package main

import (
	"fmt"
)

type DCLL interface {
	insert(val int)
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
type DCLinkedList struct {
	root *Node
	tail *Node
}

func (dcll *DCLinkedList) insert(val int) {
	node := &Node{value: val}
	if dcll.root == nil {
		dcll.root = node
		dcll.tail = node
		node.prior = node
		node.next = node
	} else {
		node.prior = dcll.tail
		node.next = dcll.root
		dcll.tail.next = node
		dcll.root.prior = node
		dcll.tail = node
	}
}
func (dcll *DCLinkedList) PrintListMoveForward() {

	current := dcll.root
	for current != nil {
		fmt.Println(current.value)
		current = current.next
		if current == dcll.root {
			break
		}
	}

}

func (dcll *DCLinkedList) PrintListMoveBack() {
	if dcll.tail == nil {
		fmt.Println("List is empty")
		return
	}

	current := dcll.tail
	for {
		fmt.Println(current.value)
		current = current.prior
		if current == dcll.tail {
			break
		}
	}
}

func (dcll *DCLinkedList) CountofNodes() int {
	count := 0
	current := dcll.root
	for current != nil {
		count++
		current = current.next
		if current == dcll.root {
			break
		}
	}
	return count
}

func (dcll *DCLinkedList) NegativeElementsbool() bool {
	current := dcll.root

	if current == nil {
		return false
	}
	for {
		if current.value < 0 {
			return true
		}
		current = current.next
		if current == dcll.root {
			break
		}
	}
	return false
}

func (dcll *DCLinkedList) SumNodes() int {
	sum := 0
	current := dcll.root

	for current != nil {
		sum++
		current = current.next
		if current == dcll.root {
			break
		}
	}
	return sum
}
func (dcll *DCLinkedList) SumOddIndexes() int {
	sumodd := 0
	count := dcll.CountofNodes()
	current := dcll.root

	for i := 1; i <= count; i += 2 {
		for current != nil {
			sumodd++
			current = current.next.next
			if current == dcll.root {
				break
			}
		}
	}
	return sumodd
}

func main() {
	var DCLList DCLL = &DCLinkedList{}

	DCLList.insert(1)
	DCLList.insert(2)

	fmt.Println("\nlist forward")
	DCLList.PrintListMoveForward()
	fmt.Println("\nlist backward")
	DCLList.PrintListMoveBack()

	fmt.Println("count of nodes՝", DCLList.CountofNodes())
	fmt.Println("negativ element՝", DCLList.NegativeElementsbool())
	fmt.Println("sum of nodes`", DCLList.SumNodes())
	fmt.Println("sum of odd nodes", DCLList.SumOddIndexes())

}
