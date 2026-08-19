/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow_ptr := head
	fast_ptr := head
	for fast_ptr != nil && fast_ptr.Next != nil {
		slow_ptr = slow_ptr.Next
		fast_ptr = fast_ptr.Next.Next
		if fast_ptr == slow_ptr {
			return true
		}
	}
	return false
}
