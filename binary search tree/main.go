package main

import "fmt"

type BSTree interface {
	insert(val int)
	preorderTree(node *Node)
}
type Node struct {
	right *Node
	left  *Node
	value int
}

type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) insert(val int) {
	node := &Node{value: val}

	if bst.root == nil {
		bst.root = node
		return
	}
	current := bst.root
	for {
		if val < current.value {
			if current.left == nil {
				current.left = node
				return

			}
			current = current.left

		} else {
			if current.right == nil {
				current.right = node
				return
			}
			current = current.right
		}
	}
}
func (bst *BinarySearchTree) preorderTree(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.value)
	bst.preorderTree(node.left)
	bst.preorderTree(node.right)

}

func main() {
	var bst BSTree = &BinarySearchTree{}
	bst.insert(6)
	bst.insert(3)
	bst.insert(7)
	bst.insert(2)
	bst.insert(4)
	bst.insert(8)

	fmt.Println(" binary search tree")
	bst.preorderTree(bst.(*BinarySearchTree).root)
}
