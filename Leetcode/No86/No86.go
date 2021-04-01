package No86

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList() *ListNode {
	var n1 ListNode = ListNode{
		Val:  1,
		Next: nil,
	}
	var n2 ListNode = ListNode{
		Val:  4,
		Next: nil,
	}
	var n3 ListNode = ListNode{
		Val:  3,
		Next: nil,
	}
	var n4 ListNode = ListNode{
		Val:  2,
		Next: nil,
	}
	var n5 ListNode = ListNode{
		Val:  5,
		Next: nil,
	}
	var n6 ListNode = ListNode{
		Val:  2,
		Next: nil,
	}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5
	n5.Next = &n6
	return &n1
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var Myhead ListNode = ListNode{
		Val:  0,
		Next: head,
	}
	p1 := &Myhead
	for p1.Next.Val < x {
		p1 = p1.Next
		if p1.Next == nil {
			return Myhead.Next
		}
	}
	pre := p1
	cur := p1.Next
	for cur != nil {
		if cur.Val < x {
			temp := p1.Next
			p1.Next = cur
			pre.Next = cur.Next
			cur.Next = temp
			cur = pre.Next
			p1 = p1.Next
		} else {
			pre = pre.Next
			cur = cur.Next
		}
	}
	return Myhead.Next
}
