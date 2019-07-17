/*
如何找出旋转数组的最小元素
把一个有序数组最开始的若干元素搬到数组的末尾，称为数组的旋转。输入一个排好序的数组的一个旋转，输出旋转数组的最小元素。例如数组{3,4,5,1,2}为数组{1,2,3,4,5}的一个旋转，该数组的最小值为1。

二分查找：

给定数组arr，首先定义两个变量low和high，分别表示数组的第一个元素和最后一个元素的下标。按照题目中对旋转规则的定义，第一个元素应该是大于或者等于最后一个元素的（当旋转个数为0，
即没有旋转的时候，要单独处理，直接返回数组第一个元素）。接着遍历数组中间的元素arr[mid]，其中mid=(high+low)/2
 */
package main

import (
	"fmt"
)

func getMinPara(arr []int,low,high int) int  {
	//如果旋转个数为0，即没有旋转，单独处理，直接返回数组头元素
	if high < low {
		return arr[0]
	}

	//只剩下一个元素一定是最小值
	if high == low {
		return arr[low]
	}

	//mid = (low + high)/2 采用下面写法防止溢出
	mid := low + ((high - low) >> 1) //右移
	//判断是否arr[mid]为最小值
	if arr[mid] < arr[mid - 1] {
		return arr[mid]
	}else if arr[mid + 1] < arr[mid]{
	    return 	arr[mid + 1]
	}else if arr[high] > arr[mid]{
		//最小值一定在数组左半部分
		return getMinPara(arr,low,mid-1)
	}else if(arr[mid] > arr[low]){
		//最小值一定在数组右半部分
		return getMinPara(arr,mid+1,high)
	}else{
		//arr[low] == arr[mid] && arr[high] == arr[mid]
		//这种情况下无法确定最小值所在的位置，需要在左右两部分分别进行查找
		left := getMinPara(arr,low,mid-1)
		right := getMinPara(arr,mid+1,high)
		if left < high {
			return left
		}else{
			return right
		}
	}
}

func getMin(arr []int) int  {
	if arr == nil {
		return -1
	}else{
		return getMinPara(arr,0,len(arr) -1)
	}
}

func swap(arr []int,low,high int)  {
	//交换数组low到high的内容
	for ;low<high ;low,high=low+1,high-1  {
		tmp :=  arr[low]
		arr[low] = arr[high]
		arr[high] = tmp
	}
}

func main()  {
	fmt.Println("找出最小值")
	arr := []int{5,6,1,2,3,4}
	fmt.Println(getMin(arr))
	arr = []int{1,1,0,1}
	fmt.Println(getMin(arr))
}
