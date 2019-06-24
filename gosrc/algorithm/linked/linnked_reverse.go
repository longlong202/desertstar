//实现链表的逆序
//给定一个带头结点的单链表，请将其逆序。即如果单链表原来为head->1->2->3->4->5->6->7，则·逆序后变为head->7->6->5->4->3->2->1
//单链表中每个结点的地址都存储在其前驱结点的指针域中，因此，对单链表中任何一个结点的访问，只能从链表的头指针开始遍历。

//next是主线的指针，pre和cur是新建的指针
/*
	next next.Next
pre cur
 */
package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

//带头结点的逆序
func Reverse(node *LNode)  {
	if node == nil || node.Next == nil {
		return
	}

	var pre *LNode //定义前驱结点
	var cur *LNode //定义当前结点
	next := node.Next //后继结点
	for next != nil {
		cur = next.Next //保存当前结点的下一个结点，防止丢失，因为下一步要将当前结点指向前一个结点，当前结点往后移动
		next.Next = pre //当前结点的下一个结点指向前一个结点
		pre = next //pre指针指向
		next = cur //当前结点成为后继结点
	}
	node.Next = pre
}

//方法二，递归实现
func RecursiveReverseChild(node *LNode) *LNode  {
	if node == nil || node.Next == nil {//判断是否为最后一次递归
		return node
	}
	newHead := RecursiveReverseChild(node.Next)//递归下一个 最后一遍4->5时，新的head即5，node.Next调用的时候返回的是node
	//1-2-3-4-5 最后一次遍历是node 4
	//4->5 转换成5->4
	node.Next.Next = node //node此时为4,则node->next为5，node->next->next 即5->4,然后4->nil,即传入4->5，回传5->4->nil
	node.Next = nil
	return newHead
}

func RecursiveReverse(node *LNode)  {
	firstNode := node.Next //从1开始递归
	//递归调用
	newHead := RecursiveReverseChild(firstNode)
	node.Next = newHead
}

//方法三 插入法
//只需要主线的next指针和cur当前指针即可
//head->1->2->3->4
//从2开始处理，即初始化的时候,cur代表当前要处理的元素,所以cur = node.Next.Next
//因为结尾最终1是尾指针，所以node.Next.Next = nil
/*
next
cur
要将head->1->2->3->4 第一步处理之后变成head->2->1 3->4
将2->1之前，需要将2->3的地址3记录下来，cur是2，则next是3，用next指针记录下一步轮询要处理的指针地址，记录好2->3的指针后，可以断开2->3，变成2->1
 */
func InsertReverse( node *LNode)  {
	if node == nil || node.Next ==nil {
		return
	}

	var cur *LNode //当前看结点指针
	var next *LNode //后继结点指针
	cur = node.Next.Next

	//设置链表第一个结点为尾结点
	node.Next.Next = nil //1的后继指针将是nil
	//把遍历到的结点插入到头结点后面
	for cur != nil {
		next = cur.Next //记录next指针，下一次轮询用
		cur.Next = node.Next //将2->3的指针变成2->1
		node.Next = cur //将head->1的指针变成 head->2 插入头结点的后面
		cur = next //将当前指针移到next指针准备下一次轮询
	}

}


func main()  {
	head := &LNode{}
	fmt.Println("就地逆序")
	CreateNode(head,8)
	PrintNode("逆序前：",head)
	//Reverse(head) //方法一
	//RecursiveReverse(head)//方法二
	InsertReverse(head)//方法三
	PrintNode("逆序后:",head)
}
