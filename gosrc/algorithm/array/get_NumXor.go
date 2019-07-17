/*
给定一个由n-1个整数组成的未排序的数组序列，其元素都是1～n中的不同的整数，请写出一个寻找数组序列中缺失整数的线性时间算法。

异或法：
在进行异或运算时，当参与运算的两个数相同时，异或结果为假，当参与异或运算的两个数不相同时，异或结果为真。
 */
package main

import "fmt"

func getNumXor(arr []int) int  {
	if arr == nil || len(arr) == 0 {
		return -1
	}

	a := arr[0]
	b :=1
	len := len(arr)
	for i:=1;i<len;i++ {
		a ^=arr[i]
	}

	for i:=2;i<=len + 1;i++ {
		b ^=i
	}
	return a^b
}

func main()  {
	fmt.Println("累加求和")
	arr := []int{1,4,3,2,7,5,6,9}
	fmt.Println(getNumXor(arr))

}
