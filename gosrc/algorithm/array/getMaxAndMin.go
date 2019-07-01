/*
给定数组a1,a2,……,an,要求找出数组中的最大值和最小值。假设数组中的值两两各不相同.

1、分治法
将数组两两一对分组，如果数组元素个数为基数个，就把最后一个元素单独分为一组，然后分别对每一组中相邻的两个元素进行比较。
把两个数中，值小的的数放在数组的左边，值大的数放在数组右边，只需要比较n/2次就可以将数组分组完成。

2、扩展的分治法
将数组分为
 */
package main

import "fmt"

//分治法
func GetMaxAndMin(arr []int) (max,min int)  {
	if arr == nil {
		return 0,0
	}

	len := len(arr)
	max = arr[0]//初始值都一样
	min = arr[0]//初始值都一样
	//两两分组,较小的放左半部，较大的放右半部
	for i:=0;i<len -1;i=i+2{
		if arr[i] > arr[i+1] {
			tmp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = tmp
		}
	}

	//在各个分组的左半部分找最小值
	for i:=2;i<len ;i = i+2  {
		if arr[i] < min{
			min = arr[i]
		}
	}

	//在各个分组的右半部分找最大值
	for i:=3;i<len ;i=i+2  {
		if arr[i] > max {
			max = arr[i]
		}
	}

	//如果数组中元素个数是奇数个，最后一个元素被分为一组，需要特殊处理
	if len % 2 == 1 {
		if max < arr[len -1] {
			max = arr[len -1]
		}

		if min > arr[len -1]{
			min = arr[len - 1]
		}
	}

	return
}

//扩展的分治法


func main()  {
	arr := []int{7,3,19,40,4,7,1}
	max,min := GetMaxAndMin(arr)
	fmt.Println("分治法")
	fmt.Println("max=",max)
	fmt.Println("min=",min)
}


