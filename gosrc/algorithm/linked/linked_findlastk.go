/*
如何找出单链表中的倒数第k个元素，例如给定单链表：1->2->3->4->5->6->7,则单链表的倒数第k=3个元素为5

快慢指针法
设置两个指针，让其中一个指针比另一个指针先前移动k步，然后两个指针同时往前移动。循环到先行指针值为nil时，另一个指针所指的位置就是所要找的位置。
 */
package main

import (
	"fmt"
	. "github.com/isdamir/gotype")

func FindLastK(head *LNode,k int) *LNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head.Next
	fast := head.Next
	i := 0
	for i=0;i<k && fast !=nil;i ++  {
		fast = fast.Next
	}

	if i < k {
		return nil
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

func main()  {
	fmt.Println("寻找倒数k")
	head := &LNode{}
	CreateNode(head,8)
	PrintNode("链表：",head)
	fmt.Println("链表倒数第3个元素为：",FindLastK(head,3).Data)
}