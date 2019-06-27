/*
数字1～1000放在含有1001个元素的数组中，其中只有唯一的一个元素值重复，其他数字均只出现一次。设计一个算法，将重复元素找出来，要求每个数组元素只能访问
一次。如果不使用辅助空间，能否设计一个算法实现？
 */
package main

import (
	"fmt"
)

func main()  {
	arr := []int{1,3,4,2,5,3}
	fmt.Println("异或法")
	fmt.Println(FindDupByXOR(arr))

}

//(1^3^4^2^5^3)^(1^2^3^4^5) = (1^1)^(2^2)^(3^3^3)^(4^4)^(5^5)=0^0^3^0^0=3
func FindDupByXOR(arr []int) (int) {
	if arr == nil {
		return -1
	}

	result :=0 //0和任何值异或都是该值
	len := len(arr)
	for _,v := range arr{
		result ^= v
	}

	for i:=1;i<len;i++{
		result ^=i
	}
	return result
}
