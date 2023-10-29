package main

import (
	"strconv"
)

func main() {
	add("243", "564")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	listOneStr := ""
	listTwoStr := ""

	for l1 != nil {
		listOneStr = strconv.Itoa(l1.Val) + listOneStr
		l1 = l1.Next
	}

	for l2 != nil {
		listTwoStr = strconv.Itoa(l2.Val) + listTwoStr
		l2 = l2.Next
	}

	list := add(listOneStr, listTwoStr)

	var head, current *ListNode

	for _, v := range list {
		newNode := &ListNode{
			Val:  v,
			Next: nil,
		}

		if head == nil {
			head = newNode
			current = head
		} else {
			current.Next = newNode
			current = newNode
		}
	}

	return head
}

func add(str1, str2 string) []int {
	int1, _ := strconv.Atoi(str1)

	int2, _ := strconv.Atoi(str2)

	total := int1 + int2
	if total == 0 {
		return []int{0}
	}

	intListReversed := []int{}

	for total > 0 {
		modulus := total % 10
		intListReversed = append(intListReversed, modulus)
		total = total / 10
	}
	return intListReversed
}

//this works but fails on leetcode for [1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1] and [5,6,4]
//the AtoI cant convert this large of a numberm gives an error here and then 9223372036854775807 as the conversion somehow
