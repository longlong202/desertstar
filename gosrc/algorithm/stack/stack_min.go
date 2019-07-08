/*
如果用O(1)的时间复杂度求栈中最小元素

实现思路：
如果当前入栈的元素比原来栈中的最小值还小，则把这个值压入保存最小元素的栈中，在出栈的时候，如果当前出栈的元素恰好为当前栈中的最小值，则保存最小值的
栈顶元素也出栈，使得当前最小值变为当前最小值入栈之前的那个最小值。
 */
package main

import (
	"fmt"
	"math"
	t "github.com/isdamir/gotype"
)

type MinStack struct {
	//用来存储栈中元素
	elemStack *t.SliceStack
	//栈顶永远存储当前elemstack中最小的值
	minstack *t.SliceStack
}

func (p *MinStack) Push(data int)  {
	p.elemStack.Push(data)
	//更新保存最小元素的栈
	if p.minstack.IsEmpty(){
		p.minstack.Push(data)
	}else{
		if data <= p.minstack.Top().(int) {
			p.minstack.Push(data)
		}
	}
}

func (p *MinStack) Pop() int  {
	topData := p.elemStack.Pop().(int)
	if topData == p.Min() {
		p.minstack.Pop()
	}

	return topData
}
func (p *MinStack) Min() int  {
	if p.minstack.IsEmpty() {
		return math.MaxInt32
	}else{
		return p.minstack.Top().(int)
	}
}

func main()  {
	stack := &MinStack{elemStack:&t.SliceStack{},minstack:&t.SliceStack{}}
	stack.Push(5)
	fmt.Println("栈中最小值为,",stack.Min())
	stack.Push(6)
	fmt.Println("栈中最小值为,",stack.Min())
	stack.Push(2)
	fmt.Println("栈中最小值为,",stack.Min())
	stack.Pop()
	fmt.Println("栈中最小值为,",stack.Min())

}

