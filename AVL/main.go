package main

import "fmt"

type AVL interface {
	insert(val int)
	insertNode(root *Node, val int) *Node
	height(root *Node) int
	rightRotate(x *Node, val int) *Node
	leftRotate(x *Node, val int) *Node
	leftrightRotate(x *Node, val int) *Node
	rightleftRotate(x *Node, val int) *Node
	preorderTraversal(node *Node)
}
type Node struct {
	right *Node
	left  *Node
	value int
}
type AVl struct {
	root *Node
}

func (avl *AVl) insertNode(root *Node, val int) *Node {
	if root == nil {
		return &Node{value: val}
	}

	if val < root.value {
		root.left = avl.insertNode(root.left, val)
	} else if val > root.value {
		root.right = avl.insertNode(root.right, val)
	} else {
		return root
	}

	balance := avl.height(root.left) - avl.height(root.right)

	if balance > 1 && val < root.left.value {
		return avl.rightRotate(root, val)
	}

	if balance < -1 && val > root.right.value {
		return avl.leftRotate(root, val)
	}

	if balance < -1 && val > root.right.value {
		return avl.leftRotate(root, val)
	}
	if balance < -1 && val < root.right.value {
		return avl.rightleftRotate(root, val)
	}

	return root
}

func (avl *AVl) insert(val int) {
	avl.root = avl.insertNode(avl.root, val)
}

func (avl *AVl) height(root *Node) int {
	if root == nil {
		return 0
	}
	leftHeight := avl.height(root.left)
	rightHeight := avl.height(root.right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (avl *AVl) rightRotate(x *Node, val int) *Node {
	height := avl.height(x)

	if height > -1 && val < x.left.value {
		y := x.left
		z := y.right

		y.right = x
		x.left = z

		return y

	}
	return x
}

func (avl *AVl) leftRotate(x *Node, val int) *Node {
	height := avl.height(x)

	if height > 1 && val > x.right.value {
		y := x.right
		z := y.left

		y.left = x
		x.right = z

		return y
	}
	return x
}
func (avl *AVl) leftrightRotate(x *Node, val int) *Node {
	height := avl.height(x)

	if height > 1 && val > x.left.value {
		x.left = avl.leftRotate(x.left, val)
		x = avl.rightRotate(x, val)

	}
	return x
}

func (avl *AVl) rightleftRotate(x *Node, val int) *Node {
	height := avl.height(x)

	if height < -1 && val < x.right.value {
		x.right = avl.rightRotate(x.right, val)
		x = avl.leftRotate(x, val)

	}
	return x
}

func (avl *AVl) preorderTraversal(node *Node) {
	if node == nil {
		return
	}

	fmt.Print(node.value, " ")
	avl.preorderTraversal(node.left)
	avl.preorderTraversal(node.right)
}

func main() {
	var Avl AVL = &AVl{}

	Avl.insert(6)
	Avl.insert(3)
	Avl.insert(7)
	Avl.insert(2)
	Avl.insert(4)
	Avl.insert(8)
	Avl.insert(9)

	fmt.Println("AVL")
	Avl.preorderTraversal(Avl.(*AVl).root)
}
