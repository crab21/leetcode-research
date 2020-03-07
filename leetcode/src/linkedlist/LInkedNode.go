package linkedlist

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type ElementType interface {}
type LinkNode struct {
	Data ElementType
	Next *LinkNode
}
type LinkedList struct {
	Head *LinkNode
}

func NewNode(data ElementType,next *LinkNode)*LinkNode{
	return &LinkNode{data,next}
}

func NewList() *LinkedList{
	head :=&LinkNode{0, nil}
	return &LinkedList{Head:head}
}

//处理链表头节点的数据
func (list *LinkedList)sizeListInc() {
	v:=reflect.ValueOf((*list.Head).Data).Int()
	list.Head.Data = v+1
}

func (list *LinkedList)sizeListDec() {
	v:=reflect.ValueOf((*list.Head).Data).Int()
	if v==0{
		return
	}
	list.Head.Data = v-1
}

//在链表的后面添加节点
func (list *LinkedList)Append(node *LinkNode) {

	if node == nil {
		log.Panic("待添加的节点不能是nil节点")
	}
	current:=list.Head
	if current.Next == nil {
		current.Next = node
		list.sizeListInc()
		return
	}
	for current.Next!= nil {
		current = current.Next
	}
	current.Next = node
	list.sizeListInc()
}
func (list *LinkedList)PreAppend(node *LinkNode) {
	if node == nil {
		log.Panic("待添加的节点不能是nil节点")
	}

	//todo
	current := list.Head
	node.Next = current.Next
	current.Next = node
	list.sizeListInc()

}

//打印节点
func (list *LinkedList)Print() {
	if list.Head.Next == nil {
		fmt.Print("空链表")
		return
	}

	//第一个节点
	_=list.Head.Next
	/*for i:=1;current!= nil;i++ {
		fmt.Printf("Node %d:    Data = %v, current node's address = %p, Next = %p\n", current.Data, current, current.Next)
		current = current.Next
	}*/
	fmt.Print("遍历完成")
}

func (list *LinkedList)Find(elementType ElementType)(*LinkNode, bool) {
	if list.Head.Next== nil {
		log.Println("链表为空")
		return nil, false
	}
	current:=list.Head
	for current.Next!=nil {
		current =current.Next
		if current.Data==elementType {
			return current,true
		}
	}
	return nil,false
}

//删除指定节点
func (list *LinkedList)Remove(elementType ElementType)error {
	current:=list.Head
	for current.Next!= nil{
		if current.Next.Data == elementType {
			current.Next = current.Next.Next
			list.sizeListDec()
			return nil
		}
		current =current.Next
	}
	return errors.New(fmt.Sprintf("Not find the node that value is %v, fail in Removing node", elementType))
}

func TTTMall(){
	list := NewList()

	fmt.Println("\n开始测试在链表后面追加节点")
	var node *LinkNode = &LinkNode{}
	for i := 1; i < 6; i++ {
		node = &LinkNode{Data: fmt.Sprintf("node%d", i)}
		list.PreAppend(node)
	}
	list.Print()

	fmt.Println("\n开始测试在链表前面插入节点")
	node = &LinkNode{Data:"node5"}
	list.PreAppend(node)
	list.Print()

	fmt.Println("\n开始测试查找链表节点")
	if linkNodePointer, ok := list.Find("node3");  ok != false {
		fmt.Printf("find node:    Data = %v, node's address = %p, Next = %p\n", linkNodePointer.Data, linkNodePointer, linkNodePointer.Next)
	}

	fmt.Println("\n开始测试删除节点")
	if err := list.Remove("node"); err != nil {
		fmt.Println("删除节点失败")
		list.Print()
	} else {
		fmt.Println("成功删除指定节点")
		list.Print()
	}
}