package single

import (
	"goleetcode/datastruct/singlelinked"
)

// Swap pairs
func SwapPairs(head *singlelinked.ListNode) *singlelinked.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	remains := head.Next.Next
	newHead := head.Next
	newHead.Next = head
	head.Next = SwapPairs(remains)

	return newHead
}
