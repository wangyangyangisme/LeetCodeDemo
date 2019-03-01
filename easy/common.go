package easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	for i, p := 0, l; p != nil; i, p = i+1, p.Next {
		if i > 0 {
			fmt.Printf("->")
		}
		fmt.Printf("%v", p.Val)
	}
	fmt.Printf("\n")
}

func (l *ListNode) Add(val int) *ListNode {
	if l == nil {
		return &ListNode{
			Val: val,
		}
	}
	q := l
	p := l.Next
	for ; p != nil; p, q = p.Next, q.Next {
	}
	(*q).Next = &ListNode{
		Val: val,
	}
	return l
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
