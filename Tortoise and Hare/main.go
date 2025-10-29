package main

import "fmt"

type TH interface {
	insert(val int)
	CraneandRabbit()
}
type Node struct {
	next  *Node
	value int
	index *Node
}

type List struct {
	root *Node
	tail *Node
}

func (l *List) insert(val int) {
	node := &Node{value: val}
	if l.root == nil {
		l.root = node
		l.tail = node
		l.tail.next = l.root
	} else {
		l.tail.next = node
		l.tail = node
		l.tail.next = l.root
	}
}

func (l *List) CraneandRabbit() {
	rabbit := l.root
	crane := l.root

	for rabbit != nil && crane != nil {
		crane = crane.next
		rabbit = rabbit.next.next

		if crane == rabbit {
			fmt.Println("this is cyclic")

			break
		}

	}

}

func main() {
	var TortoiseRabbit TH = &List{}

	TortoiseRabbit.insert(1)
	TortoiseRabbit.insert(2)
	TortoiseRabbit.insert(3)
	TortoiseRabbit.insert(4)
	TortoiseRabbit.insert(5)
	TortoiseRabbit.insert(6)

	TortoiseRabbit.CraneandRabbit()

}
