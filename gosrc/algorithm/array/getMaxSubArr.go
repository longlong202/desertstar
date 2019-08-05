/*
如何求数组连续最大和？
一个有n个元素的数组，这n个元素既可以是正数也可以是负数，数组中连续的一个或多个元素可以组成一个连续的子数组，一个数组可能有多个这种连续的子数组，
求子数组和的最大值。例如，对于数组{1,-2,4,8,-4,7,-1,-5}而言，其最大和的子数组为{4,8,-4,7}，最大值为15


 */
package main

import (
	"fmt"
	"math"
)

/*
方法一：蛮力法
时间复杂度O(n^3)，效率太低
*/
func maxSubArray(arr []int) int{
	if arr == nil || len(arr) < 1 {
		return -1
	}

	MaxSum := 0
	len := len(arr)
	i:=0
	j:=0
	k:=0
	for i=0;i<len;i++ {
		for j=i;j<len;j++ {
			ThisSum := 0
			for k=i;k<j;k++ {
				ThisSum += arr[k]
			}

			if ThisSum  > MaxSum{
				MaxSum = ThisSum
			}
		}
	}
	return MaxSum
}

/*
方法二：重复利用已经计算的子数组和
由于Sum[i,j] = Sum[i,j-1] + arr[j] ，在计算Sum[i,j]的时候可以使用前面已算出的Sum[i,j-1]而不需要重新计算，采用这种方法可以省去计算Sum[i,j-1]的时间
时间复杂度O(n^2)
 */

func maxSubArraySum(arr []int) int  {
	if arr == nil || len(arr) < 1 {
		return -1
	}
	len := len(arr)
	maxSum := math.MinInt64
	for i:=0;i<len;i++{
		sum := 0
		for j:=i;j<len;j++ {
			sum += arr[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

/*
方法三：动态规划法
首先可以根据数组最后一个元素arr[n-1]与最大子数组的关系分为以下三种情况：
1、最大子数组包含arr[n-1]，即最大子数组以arr[n-1]结尾
2、arr[n-1]单独构成最大子数字
3、最大子数组不包含arr[n-1]，那么求arr[1...n-1]的最大子数组可以转换成求arr[1...n-2]的最大子数组

假设已经计算出子数组arr[1...i-2]的最大子数组和all[i-2]，同时也计算出arr[0...i-1]中包含arr[i-1]的最大子数组和为End[i-1]，则可以得出以下关系
All[i-1] = max{End[i-1],arr[i-1],All[i-2]}
时间复杂度和空间复杂度都为O(n)
 */

func maxSubArrayDyn(arr []int) int  {
	if arr == nil || len(arr) < 1 {
		return -1
	}

	n := len(arr)
	End := make([]int,n)
	All := make([]int,n)
	End[n-1] = arr[n-1]
	All[n-1] = arr[n-1]
	End[0] = arr[0]
	All[0] = arr[0]
	for i:=1; i<n ; i++  {
		End[i] = int(math.Max(float64(End[i-1]+arr[i]),float64(arr[i])))
		All[i] = int(math.Max(float64(End[i]),float64(All[i-1])))
	}
	return All[n-1]
}

/*
优化的动态规划：
 */
func maxSubArrayDyn2(arr []int) int  {
	if arr == nil || len(arr) < 1 {
		return -1
	}

	//最大子数组和
	nAll := arr[0]
	//包含最后一个元素的最大子数组和

	nEnd := arr[0]
	for _,v := range arr {
		nEnd = int(math.Max(float64(nEnd + v),float64(v)))
		nAll = int(math.Max(float64(nEnd),float64(nAll)))
	}
	return nAll
}

func main()  {
	arr := []int{1,-2,4,8,-4,7,-1,-5}
	fmt.Println("蛮力法")
	fmt.Println("连续最大和为：",maxSubArray(arr))

	fmt.Println("重复数值法")
	fmt.Println("连续最大和为：",maxSubArraySum(arr))

	fmt.Println("动态规划法")
	fmt.Println("连续最大和为：",maxSubArrayDyn(arr))

	fmt.Println("动态规划法优化")
	fmt.Println("连续最大和为：",maxSubArrayDyn2(arr))
}
