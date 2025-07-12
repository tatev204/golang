package main

import (
	"fmt"
)

type CLL interface {
	insert(val int)
	PrintList()
	CountofNodes() int
	NegativeElementsbool() bool
	SumNodes() int
	SumOddIndexes() int
}
type Node struct {
	value int
	next  *Node
}

type CLinkedList struct {
	root *Node
	tail *Node
}

func (cll *CLinkedList) insert(val int) {
	node := &Node{value: val}
	if cll.root == nil {
		cll.root = node
		cll.tail = node
		cll.tail.next = cll.root
	} else {
		cll.tail.next = node
		cll.tail = node
		cll.tail.next = cll.root
	}
}

func (cll *CLinkedList) PrintList() {

	current := cll.root
	for current != nil {
		fmt.Println(current.value)
		current = current.next
		if current == cll.root {
			break
		}
	}

}

func (cll *CLinkedList) CountofNodes() int {
	count := 0
	current := cll.root
	for current != nil {
		count++
		current = current.next
		if current == cll.root {
			break
		}
	}
	return count
}

func (cll *CLinkedList) NegativeElementsbool() bool {
	current := cll.root

	if current == nil {
		return false
	}
	for {
		if current.value < 0 {
			return true
		}
		current = current.next
		if current == cll.root {
			break
		}
	}
	return false
}

func (cll *CLinkedList) SumNodes() int {
	sum := 0
	current := cll.root

	for current != nil {
		sum++
		current = current.next
		if current == cll.root {
			break
		}
	}
	return sum
}

func (cll *CLinkedList) SumOddIndexes() int {
	sumodd := 0
	count := cll.CountofNodes()
	current := cll.root

	for i := 1; i <= count; i += 2 {
		for current != nil {
			sumodd++
			current = current.next.next
			if current == cll.root {
				break
			}
		}
	}
	return sumodd
}

func main() {
	var CLList CLL = &CLinkedList{}

	CLList.insert(1)
	CLList.insert(2)

	fmt.Println("\nlist ")
	CLList.PrintList()
	fmt.Println("count of nodes՝", CLList.CountofNodes())
	fmt.Println("negativ element՝", CLList.NegativeElementsbool())
	fmt.Println("sum of nodes`", CLList.SumNodes())
	fmt.Println("sum of odd nodes", CLList.SumOddIndexes())

}
