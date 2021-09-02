package linkedlist

import "fmt"

type LinkedList struct {
	Head   *Node
	Length int
}

type Node struct {
	Next *Node
	Data interface{}
}

func (l *LinkedList) Prepend(n *Node) {
	current := l.Head

	l.Head = n
	l.Head.Next = current

	l.Length++
}

func (l *LinkedList) Append(n *Node) {
	if l.Head == nil {
		l.Head = n

		l.Length++

		return
	}

	lastNode := l.Head

	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}

	lastNode.Next = n

	l.Length++
}

func (l *LinkedList) Delete(value interface{}) {
	if l.Length == 0 {
		return
	}

	if l.Head.Data == value {
		l.Head = l.Head.Next

		l.Length--

		return
	}

	currentNode := l.Head

	for currentNode.Next != nil {
		if currentNode.Next.Data == value {
			currentNode.Next = currentNode.Next.Next

			l.Length--

			return
		}

		currentNode = currentNode.Next
	}
}

func (l LinkedList) PrintData() {
	if l.Length == 0 {
		fmt.Println("List is empty!")

		return
	}

	current := l.Head

	fmt.Println("Node data:")

	for current != nil {
		fmt.Printf("%v", current.Data)

		if current.Next != nil {
			fmt.Print(" -> ")
		}

		current = current.Next
	}

	fmt.Println()
}
