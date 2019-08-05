/*
如何求集合的所有子集？
有一个集合，求其全部子集（包含集合自身）。给定一个集合s，它包含两个元素<a,b>,则其全部子集为<a,ab,b>

方法一（位图法）：
设原集合为<a,b,c,d>，数组A的某次"加1"后的状态为[1,0,1,1]，则本次输出的子集为<a,c,d>，算法的重点是模拟数组加1的操作。数组可以以一直加1，直到数组内所有元素都是1。
缺点是如果数组中有重复数时，这种方法会得到重复子集
 */
package main

import "fmt"

func getAllSySet(arr []rune,mask []int,c int){
	len := len(arr)

	if len == c {
		fmt.Print("{")
		for i,v := range arr {
			if mask[i] == 1 {
				fmt.Print(string(v)," ")
			}
		}
		fmt.Println("}")
	}else {
		mask[c] =1
		getAllSySet(arr,mask,c+1)
		mask[c]=0
		getAllSySet(arr,mask,c+1)

	}
}

func main()  {
	arr := []rune{'a','b','c'}
	mask := []int{0,0,0}
	fmt.Println("位图法")
	getAllSySet(arr,mask,0)
}
