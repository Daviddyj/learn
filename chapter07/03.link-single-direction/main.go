package main

import "fmt"

type linkNode struct {
	data int
	next *linkNode
}

func main() {
	n1 := &linkNode{
		data: 1,
	}
	n2 := &linkNode{
		data: 2,
	}
	n3 := &linkNode{
		data: 3,
	}
	n4 := &linkNode{
		data: 4,
	}
	n6 := &linkNode{
		data: 6,
	}
	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n6

	n5 := &linkNode{data: 5}
	insertLink(n1, n5)

	//n10 := &linkNode{data: 1}
	//n11 := &linkNode{data: 5}
	//insertLink(n1, n10)
	//insertLink(n1, n11)

	//rangeLink(n1)

	//deleteLink(n1, n11)
	deleteLink(n1, n6)

	a := deleteLink(n1, n1) //在删除头节点的时候需要return接受一个新的头节点
	rangeLink(a)

}

func rangeLink(head *linkNode) {
	tmpNode := head
	for {
		if tmpNode == nil {
			break
		}
		fmt.Println(tmpNode.data)
		tmpNode = tmpNode.next

	}
}

func insertLink(head *linkNode, newNode *linkNode) {
	tmpNode := head
	for {
		if tmpNode.next == nil {
			tmpNode.next = newNode
			break
		}
		if newNode.data >= tmpNode.data && newNode.data <= tmpNode.next.data {
			//插入数据
			newNode.next = tmpNode.next
			tmpNode.next = newNode
			break
		}
		tmpNode = tmpNode.next
	}

}

func deleteLink(head *linkNode, deleteNode *linkNode) *linkNode {
	tmpNode := head
	for {
		if tmpNode.next.data == deleteNode.data {
			tmpNode.next = deleteNode.next
			deleteNode.next = nil
			return head
		}
		if head == deleteNode {
			head = head.next
			return head
		}
		tmpNode = tmpNode.next
	}
}
