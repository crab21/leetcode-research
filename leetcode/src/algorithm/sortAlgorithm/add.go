package sortAlgorithm

import (
	"encoding/json"
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeListNode() {
	v1 := ListNode{5, nil}
	v1.Next = &ListNode{4, nil}
	v1.Next.Next = &ListNode{3, nil}

	v2 := ListNode{5, nil}
	v2.Next = &ListNode{6, nil}
	v2.Next.Next = &ListNode{4, nil}

	ss := addTwoNumbers(&v1, &v2)

	sa, _ := json.Marshal(*ss)
	fmt.Println(string(sa))

}

func nodeNextAdd(l1 *ListNode, l2 *ListNode, add ListNode, vp int, flag int) *ListNode {
	if l1 == nil && l2 == nil && flag == 0 {
		return nil
	}

	if l1 != nil {
		vp += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		vp += l2.Val
		l2 = l2.Next
	}

	vp += flag

	flag = vp / 10
	vp = vp % 10

	add.Val = vp
	add.Next = nodeNextAdd(l1, l2, add, 0, flag)
	return &add
}
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var sp ListNode
	return nodeNextAdd(l1, l2, sp, 0, 0)
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var sp ListNode = ListNode{}
	var add *ListNode = &sp

	flag := 0
	for ; l1 != nil || l2 != nil || flag != 0; {
		//var ss *ListNode
		tmp := 0
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
			//ss =l1
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
			//ss=l2
		}
		tmp += flag

		flag = tmp / 10
		tmp %= 10

		//if ss==nil{
		ss := &ListNode{}
		//}

		ss.Val = tmp

		add.Next = ss
		add = add.Next
	}
	return sp.Next
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l *ListNode = &ListNode{}
	pre := l
	carry := 0
	for l1 != nil || l2 != nil {
		pre.Next = &ListNode{}
		p := pre.Next
		x := 0
		y := 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		p.Val = (x + y + carry) % 10
		carry = (x + y + carry) / 10
		pre = p

	}
	if carry != 0 {
		pre.Next = &ListNode{Val: carry}
	}
	return l.Next
}