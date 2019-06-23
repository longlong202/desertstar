package main

import "fmt"

func adjustMinHeap(a []int,pos,llen int)  {
	var temp int
	child := 0
	for temp = a[pos];2 * pos + 1 <= llen;pos =child{
		child = 2 * pos + 1
		if child < llen && a[child] > a[child + 1]{
			child ++
		}

		if a[child] <  temp {
			a[pos] = a[child]
		}else {
			break
		}
	}
	a[pos] = temp
}

func MinHeapSort(array []int)  {
	llen := len(array)
	for i := llen/2 -1;i>=0;i--{
		adjustMinHeap(array,i,llen-1)
	}

	for i := llen -1;i > 0;i--{
		tmp := array[0]
		array[0] = array[i]
		array[i] = tmp //最后一个和第一个交换位置
		adjustMinHeap(array,0,i-1)
	}
}

func main()  {
	fmt.Printf("堆排序")
	data := []int{5,4,9,8,7,6,0,1,3,2}
	//data := []int{5,4,9,8,7,6,0,1,3,2,534,323,2,434,2,434,2,6,7}
	MinHeapSort(data)
	fmt.Println(data)
}
