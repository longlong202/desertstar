/*
使用链表实现栈的基本操作
 */
package main

import (
	"fmt"
	"github.com/isdamir/gotype"
	"errors"
)

type LinkedStack struct {
	head *gotype.LNode
}

func (p *LinkedStack) IsEmpty() bool  {
	return p.head.Next == nil
}

func (p *LinkedStack) Size() int  {
	size := 0
	node := p.head.Next
	for node != nil {
		node = node.Next
		size++
	}

	return size
}

func (p *LinkedStack) Push(e int)  {
	node := &gotype.LNode{Data:e,Next:p.head.Next}
	p.head.Next = node //将头结点指向新插入的结点
}

func (p *LinkedStack) Pop() int  {
	tmp := p.head.Next //tmp即为即将要弹出的结点
	if tmp != nil {
		p.head.Next = tmp.Next
		return tmp.Data.(int) //这样强制转换
	}

	panic(errors.New("栈已经为空"))
}

func (p *LinkedStack) Top() int  {
	if p.head.Next != nil {
		return p.head.Next.Data.(int)
	}
	panic(errors.New("栈已经为空"))
}

func main()  {
	LinkedMode()
}

func LinkedMode()  {
	defer func() {
		if err:= recover();err!=nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("链表构建栈结构")
	linkedStack := &LinkedStack{head:&gotype.LNode{}}
	linkedStack.Push(1)
	linkedStack.Push(2)
	linkedStack.Push(3)
	fmt.Println("栈顶元素为：",linkedStack.Top())
	fmt.Println("栈大小为：",linkedStack.Size())
	linkedStack.Pop()
	fmt.Println("弹栈成功：",linkedStack.Size())
	linkedStack.Pop()
}
