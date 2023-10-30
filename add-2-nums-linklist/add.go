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

	//get all the values of the list and convert to a string in reverse order
	for l1 != nil {
		listOneStr = strconv.Itoa(l1.Val) + listOneStr
		l1 = l1.Next
	}

	for l2 != nil {
		listTwoStr = strconv.Itoa(l2.Val) + listTwoStr
		l2 = l2.Next
	}


	list := add(listOneStr, listTwoStr)
	//at this point we have a int array in the correct order we want (reversed)

	//create a LL using an array
	var head, tail *ListNode

	for _, v := range list {
		newNode := &ListNode{
			Val:  v,
			Next: nil,
		}

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

func add(str1, str2 string) []int {
	int1, _ := strconv.Atoi(str1)

	int2, _ := strconv.Atoi(str2)

	total := int1 + int2
	if total == 0 {
		return []int{0}
	}

	intListReversed := []int{}
	//splitting an int to []int{} in this way will give us the int in revers order, for this use case this is what we want.
	for total > 0 {
		modulus := total % 10
		intListReversed = append(intListReversed, modulus)
		total = total / 10
	}
	return intListReversed
}

//this works but fails on leetcode for [1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1] and [5,6,4]
//the AtoI cant convert this large of a numberm gives an error here and then 9223372036854775807 as the conversion somehow
