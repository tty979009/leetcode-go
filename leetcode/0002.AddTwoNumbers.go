package leetcode

/*
problem：
	好像我自己的快一点
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbersM(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}
	return head.Next
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}
