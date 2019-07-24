/*
如何求数组中两个元素的最小距离？
给定一个数组，数组中含有重复元素，给定两个数字num1和num2，求这两个数字在数组中出现的位置的最小距离

动态规划法：
把每次遍历的结果都记录下来从而可以减少遍历次数。
1、当遇到num1时，记录下num1值对应的数组下标的位置lastPos1，通过求lastPos1与上次遍历到num2下标的位置的值lastPos2的差可以求出最近一次遍历到的num1与num2的距离；
2、当遇到num2时，同样记录它在数组中下标的位置lastPos2，然后通过求lastPos2与上次遍历到num1的下标值lastPos1，求出最近一次遍历到的num1和num2的距离
 */
package main

import (
	"fmt"
	"math"
)

func minDynDistance( arr []int,num1,num2 int) int  {
	if arr  == nil || len(arr) <=0 {
		return math.MaxInt64
	}

	lastPos1 := -1 //上次遍历到num1的位置
	lastPos2 := -1 //上次遍历到num2的位置
	minDis := math.MaxInt64 //num1与num2的最小距离
	for i,v :=range arr{
		if v == num1{
			lastPos1 = i
			if lastPos2 >= 0 {
				minDis  = int(math.Min(float64(minDis),float64( lastPos1-lastPos2)))
			}
		}
		if v == num2{
			lastPos2 = i
			if lastPos1 >=0{
				minDis  = int(math.Min(float64(minDis),float64( lastPos2-lastPos1)))
			}
		}
	}
	return minDis

}

func main()  {
	arr := []int{4,5,6,4,7,4,6,4,7,8,5,6,4,3,10,8}
	num1 := 5
	num2 := 10
	fmt.Println("动态规划法：")
	fmt.Println(minDynDistance(arr,num1,num2))
}
