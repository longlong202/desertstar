/*
如何对链表进行重新排序
给定链表L0->L1->...->Ln-1->Ln，把链表重新排序为L0->Ln->L1->Ln-1.....。要求：1、在原来链表的基础上进行排序，即不能申请新的结点；2、只能修改结点的next域，不能修改数据域。

思路：
1、找到中间结点；
2、对链表的后半部分子链表进行逆序；
3、将前半部分链表和后半部分链表进行合并
 */
package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)
//找出链表的中间结点
func findMiddleNode(head *LNode) *LNode  {
	if head == nil || head.Next == nil {
		return head
	}

	fast := head  //快指针每次走两步
	slow := head //慢指针每次走一步
	slowPre := head
	for fast != nil && fast.Next != nil { //快指针还没到尽头，则一直遍历
		slowPre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	slowPre.Next = nil //分割链表成为两部分
	return slow
}

//对不带头结点的单链表进行翻转
func reverse(head *LNode) *LNode  {
	if head == nil || head.Next == nil {
		return head
	}

	var pre *LNode //前驱结点
	var next *LNode //当前结点
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

//重新排序
func Reorder(head *LNode)  {
	if head == nil || head.Next == nil {
		return
	}
	cur1 := head.Next //前半部分链表的第一个结点
	mid := findMiddleNode(head.Next)
	cur2 := reverse(mid) //后半部分链表逆序后的的第一个结点

	var tmp *LNode

	//合并链表
	for cur1.Next !=nil  {
		tmp = cur1.Next
		cur1.Next = cur2
		cur1 = tmp
		tmp = cur2.Next
		cur2.Next = cur1
		cur2 = tmp
	}
	cur1.Next =  cur2
}

func main()  {
	fmt.Println("链表排序")
	head := &LNode{}
	CreateNode(head,8)
	PrintNode("排序前:",head)
	Reorder(head)
	PrintNode("排序后:",head)

}
