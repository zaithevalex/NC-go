package main

import "fmt"

type Node[T any] struct {
	next, prev *Node[T]
	value      T
}

func NodeCreate[T any](value T, next *Node[T], prev *Node[T]) *Node[T] {
	return &Node[T]{next, prev, value}
}

func (node *Node[T]) Add(value T) {
	var saveNode = node
	for saveNode.next != nil {
		saveNode = saveNode.next
	}
	saveNode.next = NodeCreate(value, nil, saveNode)
}

func (node *Node[T]) Remove(index int) int {
	var saveNode = node
	var counter = 0

	for saveNode.next != nil {
		if counter == index {
			if saveNode.prev != nil {
				saveNode.prev.next = saveNode.next
			}
			if saveNode.next != nil {
				saveNode.next.prev = saveNode.prev
			}
			saveNode = nil
			return index
		}
		saveNode = saveNode.next
		counter++
	}
	return -1
}

func (node *Node[T]) Len() int {
	var saveNode = node
	var counter = 0
	for saveNode.next != nil {
		saveNode = saveNode.next
		counter++
	}
	return counter
}

func Print[T any](node Node[T]) {
	fmt.Printf("elem: %v\n", node)
	if node.next == nil {
		fmt.Println()
		return
	}
	Print(*node.next)
}

func main() {
	var list = Node[int]{nil, nil, 0}
	list.Add(1)
	list.Add(2)
	list.Add(17)
	Print(list)
	list.Remove(1)
	Print(list)

	//a := 4
	//fmt.Println(a)

	//b := &a
	//(*b)++
	//fmt.Println(a)
}
