package linkedlists

//ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	ln := &ListNode{}
	ln.Next = head
	fast := ln
	slow := ln
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/
