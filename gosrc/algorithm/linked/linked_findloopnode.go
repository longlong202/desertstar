/*
如何检测一个较大的单链表是否有环？
单链表有环指的是单链表中某个结点的next域指向的是链表中在它之前的某一个结点，这样在链表的尾部形成一个环形结构。如何判断单链表是否有环存在？

一、快慢指针法
定义两个指针fast（快）与slow（慢），二者的初始值都指向链表头，指针slow每次前进一步，指针fast每次前进两步，两个指针同时向前移动，快指针每移动一次都要跟慢指针比较，如果快指针等于慢指针，
就证明这个链表是带环的单向链表，否则，证明这个链表是不带环的循环链表。
 */
package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

//判断单链表是否有环
func IsLoop(head *LNode) *LNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head.Next
	fast := head.Next

	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}

//找到环的入口点
func FindLoopNode(head *LNode,meetNode *LNode) *LNode  {
	first := head.Next
	second :=  meetNode

	for first != second {
		first = first.Next
		second = second.Next
	}

	return first

}

func main()  {
	fmt.Println("查找环")
	head := &LNode{}
	CreateNode(head,8)
	meetNode := IsLoop(head)
	if meetNode != nil{
		fmt.Println("有环")
		loopNode := FindLoopNode(head,meetNode)
		fmt.Println("环的入口点为：",loopNode)
	}else{
		fmt.Println("无环")
	}
}
