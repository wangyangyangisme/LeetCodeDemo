package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	for p, q := head, head.Next; q != nil; {
		if p.Val == q.Val {
			p.Next = q.Next
			q = q.Next
		} else {
			p, q = p.Next, q.Next
		}
	}
	return head
}
