package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	// Initialize the slow and fast pointers.
	slow := head
	fast := head

	// Traverse the list with the slow and fast pointers.
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // Move the slow pointer by one step.
		fast = fast.Next.Next // Move the fast pointer by two steps.

		// If the fast pointer meets the slow pointer, there is a cycle.
		if slow == fast {
			return true
		}
	}

	// If the fast pointer reaches the end of the list, there is no cycle.
	return false
}
