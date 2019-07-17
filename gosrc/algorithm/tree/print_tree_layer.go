/*
如何从顶部开始逐层打印二叉树结点数据？

分析：
为了实现对二叉树的层序遍历，就要求在遍历一个结点的同时记录下它的孩子结点的信息，然后按照这个记录的顺序来访问结点的数据，在实现的时候可以采用队列来存储当前
遍历到的结点的孩子结点，从而实现二叉树的层序遍历。
 */
package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

func PrintTreeLayerV2(root *BNode)  {
	if root == nil {
		return
	}

	var p *BNode
	queue := NewSliceQueue()
	//树根结点进队列
	queue.EnQueue(root)
	for queue.Size() > 0 {
		p = queue.DeQueue().(*BNode)
		//访问当前结点
		fmt.Println(p.Data," ")
		//如果这个结点的左孩子不为空则入队列
		if p.LeftChild != nil {
			queue.EnQueue(p.LeftChild)
		}

		//如果这个结点的右孩子不为空则入队列
		if p.RightChild != nil {
			queue.EnQueue(p.RightChild)
		}
	}
}

func main()  {
	data := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("数组：",data)
	root := ArrayToTree(data,0,len(data) - 1)
	fmt.Println("树的层序遍历结果为：")
	PrintTreeLayerV2(root)
	fmt.Println()
}
