/*
给定一个没有排序的链表，去掉重复项，并保留原顺序，例如链表 1->3->1->5->5->7，去掉重复项后变为 1->3->5->7
 */
package main

import (
	"fmt"
	. "github.com/isdamir/gotype")

//方法一，顺序删除
/*
外层循环用一个指针从第一个结点开始遍历整个链表，然后内层循环用另外一个指针遍历其余结点，将与外层循环遍历到的指针所指结点的数据域相同的结点删除
 */

func RemoveDup(head *LNode)  {
	if head == nil || head.Next == nil {
		return
	}

	outerCur := head.Next //外层指针
	var innerCur *LNode //内层指针
	var innerPre *LNode //innerCur的前驱指针

	for ;outerCur != nil ; outerCur = outerCur.Next  { //外层指针往后
		for innerCur,innerPre = outerCur.Next,outerCur;innerCur !=nil;{
			if outerCur.Data == innerCur.Data {
				innerPre.Next = innerCur.Next
				innerCur = innerCur.Next
			}else {
				innerPre = innerCur
				innerCur = innerCur.Next
			}
		}
	}
}

//创建链表
func CreateNodeT(node *LNode, max int) {
	cur := node
	for i := 1; i < max; i++ {
		cur.Next = &LNode{}
		num := i
		if i == 2{
			num = 3
		}
		if i == 3{
			num = 1
		}
		if i >=4 && i<=5{
			num = 5
		}
		if i > 5{
			num = 7
		}
		cur.Next.Data = num
		cur = cur.Next
	}
}

func main()  {
	head := &LNode{}
	fmt.Println("删除重复结点：")
	CreateNodeT(head,7)
	fmt.Println("顺序删除")
	PrintNode("删除重复结点前,",head)
    RemoveDup(head)
	PrintNode("删除重复结点后,",head)
}

