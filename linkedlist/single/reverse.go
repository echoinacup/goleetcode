package single

import (
	"goleetcode/datastruct/singlelinked"
)

// Revers Linked List
func ReverseList(head *singlelinked.ListNode) *singlelinked.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	reversed := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return reversed
}
