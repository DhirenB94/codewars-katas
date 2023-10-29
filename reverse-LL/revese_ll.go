package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinkedList(values []int) *ListNode {
	var head, tail *ListNode
	for _, val := range values {
		newNode := &ListNode{Val: val}
		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
	}
	return head
}

func reverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode //nil
	current := head    // 1 -> 2 -> 3 -> 4 -> 5 -> nil
	for current != nil {
		next := current.Next //2 -> 3 -> 4 -> 5 -> nil
		current.Next = prev  //nil
		prev = current       //1 -> nil
		current = next       //2 -> 3 -> 4 -> 5 -> nil
	}
	return prev
}

func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	values := []int{1, 2, 3, 4, 5}
	head := createLinkedList(values)
	fmt.Println("Original Linked List:")
	printLinkedList(head)

	reversed := reverseLinkedList(head)
	fmt.Println("Reversed Linked List:")
	printLinkedList(reversed)
}
