package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	newList := &ListNode{}
	current := newList

	//iterate through both lists
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			//add list1 node to newlist and move list1 along one node
			current.Next = list1
			list1 = list1.Next
		} else {
			//add list2 node to newlist and move list2 along one node
			current.Next = list2
			list2 = list2.Next
		}
		//move the newList node along by one
		current = current.Next
	}

	//if there are any remaining nodes in either list, add them to the new list
	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}

	return newList.Next
}

//recursive solution
// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//     if l1 == nil && l2 == nil {
//         return nil
//     }
//     if l1 != nil && l2 == nil {
//         return l1
//     }
//     if l1 == nil && l2 != nil {
//         return l2
//     }

//     if l1.Val <= l2.Val {
//         l1.Next = mergeTwoLists(l1.Next, l2)
//         return l1
//     }
//     l2.Next = mergeTwoLists(l1, l2.Next)
//     return l2
//}
