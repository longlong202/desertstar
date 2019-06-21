package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baton := make(chan int)//指挥棒

	wg.Add(1)//最后一个跑步者计数加1

	go Runner(baton)

	//开始比赛
	baton <-1

	//等待比赛结束
	wg.Wait()
}

//模拟接力赛中每一位跑步者
func Runner(baton chan int)  {
	var newRunner int

	//等待接力棒
	runner := <-baton //上一个人不传值就会阻塞在这一步

	//开始绕着跑道跑步
	fmt.Printf("Runner %d Running With Baton\n",runner)

	//创建下一位跑步者
	if runner !=4 {//如果不是最后一个跑步者
		newRunner = runner +1
		fmt.Printf("Runner %d To the Line\n",newRunner)
		go Runner(baton) //开启goroutine，但是baton值可以后面再传，会阻塞
	}

	//围绕跑道跑
	time.Sleep(100 * time.Microsecond)

	//判断比赛是否结束
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n",runner)
		wg.Done()
		return
	}

	//将接力棒传递给下一个跑步者
	fmt.Printf("Runner %d Exchange With Runner %d\n",runner,newRunner)

	baton <-newRunner //这个时候再传值即可
}