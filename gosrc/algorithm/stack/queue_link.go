/*
采用链表实现队列
分别使用两个指针指向队列的首元素与尾元素，
 */
package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
	"github.com/pkg/errors"
)

type LinkedQueueV2 struct {
	head *LNode
	end *LNode
}

//判断队列是否为空
func (p *LinkedQueueV2) IsEmpty() bool  {
	return p.head==nil
}

//获取栈中元素的个数
func (p *LinkedQueueV2) Size() int  {
	size := 0
	node := p.head
	for node != nil {
		node = node.Next
		size++
	}

	return size
}

//入队列：把元素e加到队列尾
func (p *LinkedQueueV2) EnQueue(e int)  {
	node := &LNode{Data:e}//指针未赋值
	if p.head == nil {
		p.head = node
		p.end = node
	}else{
		p.end.Next = node
		p.end = node
	}
}

//出队列：删除队列首元素
func (p *LinkedQueueV2) DeQueue()  {
	if p.head == nil {
		panic(errors.New("队列已经为空"))
	}

	p.head = p.head.Next
	if p.head == nil {
		p.end = nil
	}
}

//取得队列首元素
func (p *LinkedQueueV2) GetFront() int  {
	if p.head == nil {
		panic(errors.New("队列已经为空"))
	}
	return p.head.Data.(int)
}

//取得队列尾元素
func (p *LinkedQueueV2) GetBack() int  {
	if p.end == nil {
		panic(errors.New("队列已经为空"))
	}
	return p.end.Data.(int)
}

func main()  {
	LinkedModeV2()
}

func LinkedModeV2()  {
	defer func() {
		if err:= recover();err!=nil{
			fmt.Println(err)
		}
	}()

	fmt.Println("Linked 构建队列结构")
	linkedQueueV2 := &LinkedQueueV2{}
	linkedQueueV2.EnQueue(1)
	linkedQueueV2.EnQueue(2)
	linkedQueueV2.EnQueue(3)
	linkedQueueV2.DeQueue()
	fmt.Println("队列头元素为：",linkedQueueV2.GetFront())
	fmt.Println("队列尾元素为：",linkedQueueV2.GetBack())
	fmt.Println("队列大小为：",linkedQueueV2.Size())
}



