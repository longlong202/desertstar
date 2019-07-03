/*
如何把链表相邻元素翻转？
把链表相邻元素翻转，例如给定链表为1->2->3->4->5->6->7，则翻转后的链表变为2->1->4->3->6->5->7

方法一、就地逆序
通过调整结点指针域的指向来直接调换相邻的两个结点。如果单链表恰好有偶数个结点，那么只需要将奇偶结点对掉即可，如果链表有奇数个结点，那么只需要将除最后一个结点外的其他结点进行奇偶对掉。
 */
package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

func Reversev2(head *LNode)  {
	if head == nil || head.Next == nil {
		return
	}

	cur := head.Next
	pre := head
	var next *LNode
	for cur != nil  && cur.Next != nil{
		next = cur.Next.Next
		pre.Next = cur.Next
		cur.Next.Next = cur
		cur.Next = next
		pre = cur
		cur = next
	}
}

func main()  {
	fmt.Println("相邻元素翻转")
	head := &LNode{}
	CreateNode(head,8)
	PrintNode("顺序输出",head)
	Reversev2(head)
	PrintNode("逆序输出",head)
}