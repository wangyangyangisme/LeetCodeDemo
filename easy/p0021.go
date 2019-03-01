package easy

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := &ListNode{}
	np := head
	for ; l1 != nil && l2 != nil; np = np.Next {
		if l1.Val < l2.Val {
			np.Next = l1
			l1 = l1.Next
		} else {
			np.Next = l2
			l2 = l2.Next
		}
	}
	if l1 != nil {
		np.Next = l1
	}
	if l2 != nil {
		np.Next = l2
	}
	return head.Next
}
