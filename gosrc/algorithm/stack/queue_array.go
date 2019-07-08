/*
如何实现队列？
实现一个队列的数据结构，使其具有入队列、出队列、查看队列首尾元素、查看队列大小等功能

方法一：数组实现
用front记录队列首元素的位置，用rear来记录队列尾元素往后一个位置
 */
package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type SliceQueue struct {
	arr []int
	front int //队列头
	rear int //队列尾
}

//判断队列是否为空
func (p *SliceQueue) IsEmpty() bool  {
	return p.front == p.rear
}

//返回队列大小
func (p *SliceQueue) Size() int  {
	return p.rear-p.front
}

//返回队列首元素
func (p *SliceQueue) GetFront() int  {
	if p.IsEmpty() {
		panic(errors.New("队列已经为空"))
	}
	return p.arr[p.front]
}

//返回队列尾元素
func (p *SliceQueue) GetBack() int  {
	if p.IsEmpty() {
		panic(errors.New("队列已经为空"))
	}

	return p.arr[p.rear - 1]//数组从0计数
}

//删除队列头元素
func (p *SliceQueue) DeQueue()  {
	if p.rear >  p.front {
		p.rear--
		p.arr=p.arr[1:]
	}else{
		panic(errors.New("队列已经为空"))
	}
}

//把新元素加入到队列尾
func (p *SliceQueue) EnQueue(item int)  {
	p.arr = append(p.arr,item)
	p.rear++
}

func main()  {
	SliceModeV2()
}

func SliceModeV2()  {
	defer func() {
		if err:=recover();err!=nil {
			fmt.Println(err)
		}
	}()


	fmt.Println("构建队列结构")
	sliceQueue := &SliceQueue{arr:make([]int,0)}
	sliceQueue.EnQueue(1)
	sliceQueue.EnQueue(2)
	sliceQueue.EnQueue(3)
	fmt.Println("队列头元素为：",sliceQueue.GetFront())
	fmt.Println("队列尾元素为：",sliceQueue.GetBack())
	fmt.Println("队列大小为：",sliceQueue.Size())
}

