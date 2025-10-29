package main

import (
	"fmt"
	"hash/fnv"
)

type Node struct {
	key   string
	value int
	next  *Node
}
type HashMap struct {
	size  int
	slice []*Node
}

func hashMap(size int) *HashMap {
	slice := make([]*Node, size)
	return &HashMap{
		size:  size,
		slice: slice,
	}
}

func (hm *HashMap) NewNode(key string, value int, index int) *Node {
	return &Node{key: key, value: value, next: hm.slice[index]}
}

func (hm *HashMap) hash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32()) % hm.size
}

func (hm *HashMap) insert(key string, value int) {
	index := hm.hash(key)

	if hm.slice[index] == nil {
		node := &Node{key: key, value: value}
		hm.slice[index] = node

	} else {
		current := hm.slice[index]
		for current.next != nil {
			current = current.next
		}
		current.next = &Node{key: key, value: value}
	}
}
func (hm *HashMap) GetValue(key string) int {
	index := hm.hash(key)
	current := hm.slice[index]
	for current != nil {
		if current.key == key {
			return current.value
		}
		current = current.next
	}
	return 0

}
func (hm *HashMap) Print() {
	fmt.Println("\nHashmap:")
	for i, node := range hm.slice {
		fmt.Printf("Slice %d: ", i)
		current := node
		for current != nil {
			fmt.Printf("(%s : %d)  ", current.key, current.value)
			current = current.next
		}
		fmt.Println("nil")
	}
}

func main() {
	hm := hashMap(5)

	hm.insert("A", 1)
	hm.insert("D", 2)
	hm.insert("A", 3)
	hm.insert("D", 2)
	hm.insert("A", 5)

	fmt.Println(" Value:", hm.GetValue("D"))
	fmt.Println(" Value:", hm.GetValue("R"))

	hm.Print()
}
