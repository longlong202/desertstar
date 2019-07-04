/*
实现一个栈的数据结构，使其具有以下方法：压栈、弹栈、取栈顶元素、判断栈是否为空以及获取栈中元素个数。

方法一：
采用数组实现栈
 */
package main

import (
	"errors"
	"fmt"
)

type SliceStack struct {
	arr []int
	stackSize int
}

func (p *SliceStack) IsEmpty()  bool {
	return p.stackSize == 0
}

func (p *SliceStack) Size() int  {
	return p.stackSize
}

func (p *SliceStack) Top() int  {
	if p.IsEmpty() {
		panic(errors.New("栈已经为空"))
	}
	return p.arr[p.stackSize-1]
}

func (p *SliceStack) Pop() int  {
	if p.stackSize > 0 {
		p.stackSize--
		ret := p.arr[p.stackSize]
		p.arr=p.arr[:p.stackSize]
		return ret
	}
	panic(errors.New("栈已经为空"))
}

func (p *SliceStack) Push(t int)  {
	p.arr = append(p.arr,t)
	p.stackSize = p.stackSize + 1
}

func main()  {
	SliceMode()
}

func SliceMode()  {
	defer func() {
		if err := recover();err!=nil{
			fmt.Println(err)
		}
	}()

	fmt.Println("Slice 	构建栈结构	")
	sliceStack := &SliceStack{arr:make([]int,0)}
	sliceStack.Push(1)
	fmt.Println("栈顶元素为：",sliceStack.Top())
	fmt.Println("栈大小为：",sliceStack.Size())
	sliceStack.Pop()
	fmt.Println("弹栈成功：",sliceStack.Size())
	sliceStack.Pop()

}
