/*
输入两个整数序列，其中一个序列表示栈的push顺序，判断另一个序列有没有可能是对应的pop顺序

方法一：
使用栈来模拟入栈顺序，具体步骤如下：
1、把push序列依次入栈，直到栈顶元素等于pop序列的第一个元素，然后栈顶元素出栈，pop序列移动到第二个元素；
2、如果栈顶继续等于pop序列现在的元素，则继续出栈并pop后移；否则对push序列继续入栈。
3、如果push序列已经全部入栈，但是pop序列未全部遍历，而且栈顶元素不等于当前pop元素，那么这个序列不是一个可能的出栈序列。如果栈为空，而且pop序列也全部被遍历过，则说明
这是一个可能的pop序列。
 */
package main

import (
	"fmt"
	"github.com/isdamir/gotype"

)

func IsPopSerial(push string,pop string) bool  {
	pushLen := len(push)
	popLen := len(pop)

	if pushLen == 0 || popLen == 0 || pushLen != popLen {
		return false
	}

	pushIndex := 0
	popIndex := 0
	pushRune := []rune(push)
	popRune := []rune(pop)

	stack := gotype.NewSliceStack()
	for pushIndex < pushLen  {
		//把push序列依次入栈，直到栈顶元素等于pop序列的第一个元素
		stack.Push(pushRune[pushIndex])
		pushIndex++
		//栈顶元素出栈，pop序列移动到下一个元素
		for !stack.IsEmpty() && stack.Top() == popRune[popIndex] {
			stack.Pop()
			popIndex++
		}
	}
	//栈为空，且pop序列中元素都被遍历过
	if stack.IsEmpty() && popIndex == popLen {
		return true
	}

	return false
}

func main()  {
	push := "12345"
	pop := "32541"
	if IsPopSerial(push,pop){
		fmt.Println(pop,"是，",push,"的一个pop序列")
	}else{
		fmt.Println(pop,"不是，",push,"的一个pop序列")
	}
}