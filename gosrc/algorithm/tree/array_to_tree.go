/*
如何把一个有序整数数组放到二叉树中

分析：
如果把一个有序的整数数组放到二叉树中，那么所构造出来的二叉树必定也是一颗有序的二叉树。鉴于此，实现思路为：取数组的中间元素作为根结点，将数组分成左右两部分，
对数组的两部分用递归的方法分别构建左右子树。
 */
package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
	)

func arrayToTree(arr []int,start int,end int) *BNode  {
	var root *BNode
	if end >= start {
		root=NewBNode()
		mid := (start + end + 1)/2
		//树的根结点为数组中间的元素
		root.Data = arr[mid]
		//递归的用左半部分数组构造root的左子树
		root.LeftChild = arrayToTree(arr,start,mid-1)
		//递归的用右半部分数组构造root的右子树
		root.RightChild = arrayToTree(arr,mid+1,end)
	}
	return root
}

func main()  {
	data := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("数组：",data)
	root := arrayToTree(data,0,len(data) -1)
	fmt.Println("转换成树的中序遍历为：")
	PrintTreeMidOrder(root)
	fmt.Println()
}
