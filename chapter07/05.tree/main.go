package main

import "fmt"

type Node struct {
	root  *Node
	left  *Node
	right *Node
	data  int
}

func main() {
	root := buildTree()
	insertNode(root, &Node{data: 43})
	insertNode(root, &Node{data: 30})
	fmt.Println("插入完成")

}

func buildTree() *Node {
	n1 := &Node{data: 51}
	n2 := &Node{data: 35}
	n3 := &Node{data: 65}
	n1.left = n2
	n2.root = n1

	n1.right = n3
	n3.root = n1
	return n1
}

func insertNode(root *Node, newNode *Node) *Node {
	if root == nil {
		return newNode
	}
	if root.data == newNode.data {
		return root
	}
	if newNode.data < root.data {
		if root.left == nil {
			root.left = newNode
			newNode.root = root
		}
		insertNode(root.left, newNode)
	} else {
		if root.right == nil {
			root.right = newNode
			newNode.root = root
		}
		insertNode(root.right, newNode)
	}
	return root
}



