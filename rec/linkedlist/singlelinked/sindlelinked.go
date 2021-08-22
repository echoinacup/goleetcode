package singlelinked

import (
	"goleetcode/ds"
)

// Swap pairs
func swapPairs(head *ds.ListNode) *ds.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	remains := head.Next.Next
	newHead := head.Next
	newHead.Next = head
	head.Next = swapPairs(remains)

	return newHead
}

// Revers Linked List
func reverseList(head *ds.ListNode) *ds.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	reversed := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return reversed
}
