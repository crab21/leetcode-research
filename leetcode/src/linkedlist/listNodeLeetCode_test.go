package linkedlist

import (
	"fmt"
	"testing"
)

func TestReserveList(t *testing.T) {
	head := &ListNodeInfo{Val: 1,
		Next: &ListNodeInfo{Val: 2,
			Next: &ListNodeInfo{Val: 3,
				Next: &ListNodeInfo{4, nil}}}}
	sp := ReserveList(head)
	fmt.Println(*sp)
}

func TestListEntryNodeOfLoop(t *testing.T) {
	head := &ListNodeInfo{Val: 1, Next: nil}
	next1:=&ListNodeInfo{Val: 2, Next: nil}
	next2:=&ListNodeInfo{Val: 3, Next: nil}
	next3:=&ListNodeInfo{Val: 4, Next: nil}
	next4:=&ListNodeInfo{Val: 5, Next: nil}
	next5:=&ListNodeInfo{Val: 6, Next: nil}

	head.Next = next1
	next1.Next = next2
	next2.Next = next3
	next3.Next = next4
	next4.Next = next5
	next5.Next = next3

	sp := ListEntryNodeOfLoop(head)
	fmt.Println(sp)
}

func TestEntryNodeOfLoop(t *testing.T) {
	head := &ListNodeInfo{Val: 1, Next: nil}
	next1:=&ListNodeInfo{Val: 2, Next: nil}
	next2:=&ListNodeInfo{Val: 3, Next: nil}
	next3:=&ListNodeInfo{Val: 4, Next: nil}
	next4:=&ListNodeInfo{Val: 5, Next: nil}
	next5:=&ListNodeInfo{Val: 6, Next: nil}

	head.Next = next1
	next1.Next = next2
	next2.Next = next3
	next3.Next = next4
	next4.Next = next5
	next5.Next = next3
	loop := EntryNodeOfLoop(head)
	fmt.Println(loop)

}