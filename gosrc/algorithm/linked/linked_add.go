/*
如何计算两个单链表所代表的数之和？
给定两个单链表，链表的每个结点代表一位数，计算两个数的和。例如：输入链表(3->1->5)和链表(5->9->2),输出8->0->8,即513+295=808，注意个位数在链表头

对链表中的结点直接进行相加操作，把相加的和存储到新的链表中对应的结点中，同时还要记录结点相加后的进位。
 */
package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

func Add(h1 *LNode,h2 *LNode) *LNode {
	if h1 == nil || h1.Next == nil {
		return h2
	}
	
	if h2 == nil || h2.Next == nil {
		return h1
	}
	
	c := 0 //记录进位
	sum := 0 //记录两个结点相加的值
	p1 := h1.Next //遍历h1
	p2 := h2.Next //遍历h2
	resultHead := &LNode{} //相加后链表头结点
	p := resultHead //指向链表resultHead最后一个结点
	for p1 !=nil && p2!=nil {
		p.Next = &LNode{}
		sum = p1.Data.(int) + p2.Data.(int) + c //强制转换
		p.Next.Data = sum % 10 //余数
		c = sum / 10 //进位
		p = p.Next
		p1 = p1.Next
		p2 = p2.Next
	}

	//链表p2比p1长
	if p1 == nil {
		for p2 != nil {
			p.Next = &LNode{}
			sum = p2.Data.(int) + c
			p.Next.Data = sum % 10
			c = sum / 10 //进位
			p = p.Next
			p2 = p2.Next
		}
	}
	//链表p1比p2长
	if p2 == nil {
		for p1 != nil {
			p.Next = &LNode{}
			sum = p1.Data.(int) + c
			p.Next.Data = sum % 10
			c = sum / 10 //进位
			p = p.Next
			p1 = p1.Next
		}
	}

	if c == 1 { //判断是否还有进位
		p.Next = &LNode{}
		p.Next.Data = 1
	}
	return resultHead

}

func CreateNodeL()(l1 *LNode,l2 *LNode)  {
	l1 = &LNode{}
	l2 = &LNode{}
	cur := l1
	for i:=1;i<7 ;i++  {
		cur.Next = &LNode{}
		cur.Next.Data =i+2
		cur = cur.Next
	}

	cur = l2
	for i:=9;i>4 ;i--  {
		cur.Next = &LNode{}
		cur.Next.Data =i
		cur = cur.Next
	}
	return
}

func main()  {
	fmt.Println("链表相加")
	l1,l2 := CreateNodeL()
	PrintNode("Head1:",l1)
	PrintNode("Head2:",l2)
	addRes := Add(l1,l2)
	PrintNode("相加后:",addRes)
}



