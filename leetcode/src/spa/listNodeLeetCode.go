package spa

type ListNodeInfo struct {
	Val int
	Next *ListNodeInfo
}

func findKthToTail(k int,node *ListNodeInfo)*ListNodeInfo {
	if node == nil {
		return nil
	}
	firstPointer:=node
	secondPointer:=node
	i:=0
	for ;firstPointer!=nil ; i++ {
		firstPointer = firstPointer.Next
		if i>=k-1 {
			secondPointer = secondPointer.Next
		}
	}

	if i>=k-1 {
		return secondPointer
	}
	return nil
}

/**
链表反转
 */
func ReserveList(head *ListNodeInfo)*ListNodeInfo {
	if head == nil || head.Next == nil  {
		return head
	}
	var prev *ListNodeInfo
	cur:=head
	for ;cur!= nil ;  {
		//cur.Next, prev, cur = prev, cur, cur.Next
		tmp:=cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}




/**
链表中环的入口节点
 */
// 双指针
func ListEntryNodeOfLoop(info *ListNodeInfo) *ListNodeInfo{
	if info == nil || info.Next == nil {
		return nil
	}

	pre:=info.Next
	stepPre := info.Next.Next
	for ;pre!= stepPre ;  {
		pre = pre.Next
		stepPre = stepPre.Next.Next
	}
	stepPre = info
	for ;pre!=stepPre ;  {
		pre = pre.Next
		stepPre = stepPre.Next
	}
	return pre
}

//断链法
func EntryNodeOfLoop(head *ListNodeInfo)*ListNodeInfo {
	if head == nil || head.Next == nil{
		return nil
	}
	prev := head
	current := head.Next
	for ;current!= nil ;  {
		prev.Next = nil
		prev = current
		current = current.Next
	}
	return prev

}