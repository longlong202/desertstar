/*
给定一个整数组，如何快速地求出该数组中第k小的数。假如数组为{4,0,1,0,2,3},那么第3小的元素是1

方法一（排序法）：
首先对数组进行排序，在排序后的数组中，下标为k-1的值就是第k小的数。由于最高效的排序算法（例如快速排序）的平均时间复杂度为O(nlogn)，因此该方法的平均时间复杂度为O(nlogn)

方法二（部分排序法）：
第一次遍历从数组中找出最小的数，第二次遍历从剩下的数中找出最小的数，第k次遍历就可以从n-k+1个数中找出最小的数

方法三（类快速排序）：
将数组array[low...high] 中某一个元素（取第一个元素）作为划分依据，然后把数组划分为三部分：1.array[low...i-1]所有元素的值都小于或等于array[i] 2.array[i]
3. array[i+1 ... high]所有元素都大于array[i]

1）如果i-low == k-1，说明array[i]就是第k小的元素，那么直接返回array[i]
2）如果i-low>k-1 ，说明第k小的元素肯定在array[low ... i-1]中，那么只需要递归地在array[low ... i-1]中找第k小的元素即可。
3) 如果i-low < k-1,说明第k小的元素肯定在array[i+1...high]中，那么只需要递归地在array[i+1...high]中找出第k-(i-low)-1小的元素即可

对于数组{4,0,1,0,2,3}，找出第3小的数：
{3,0,1,0,2},{4},{}
{2,0,1,0},{3},{}
{0,0,1},{2},{}
{0},{0},{1}

 */
package main

import "fmt"

func FindSmallK(arr []int,low,high,k int)int{
	i := low
	j := high
	spliElem := arr[i]
	//把小于等于splitElem的数放到数组中splitelem的左边，大于splitElem的值放在右边
	for i < j  {
		for i<j && arr[j] >= spliElem {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i<j && arr[i] <= spliElem{
			i++
		}
		if i<j{
			arr[j] = arr[i]
			j--
		}

	}

	arr[i] = spliElem
	//splitElem在子数组array[low ~ high] 中下标的偏移量
	subArrayIndex := i-low
	//如果splitElem在array[low ~ high]所在的位置恰好为k-1，那么它就是最小第k小的元素
	if subArrayIndex == k-1 {
		return arr[i]
	}else if subArrayIndex > k-1{
		//如果splitElem在array[low ~ high]所在的位置大于k-1，那么只需在array[low ~ i-1] 中找第k小的元素
		return FindSmallK(arr,low,i-1,k)
	}else{
		//在array[i+1 ~ high]中找第k-i+low-1小的元素
		return FindSmallK(arr,i+1,high,k-(i-low)-1)
	}
}

func main()  {
	k := 3
	arr := []int{4,0,1,0,2,3}
	fmt.Println("第",k,"小的值为",FindSmallK(arr,0,len(arr)-1,k))
}
