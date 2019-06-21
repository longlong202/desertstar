//通过无缓冲通道来模拟两个2个goroutine之间的网球比赛

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init(){
	rand.Seed(time.Now().UnixNano())
}

func main()  {
	game := make(chan int) //创建无缓冲通道
	wg.Add(2) //等待2个goroutine

	//启动两个选手 TODO
	go player("Jessie",game) //传入通道
	go player("Jacob",game) //传入通道

	//发球
	game <- 1

	//等待游戏结束
	wg.Wait()
}

//模拟选手打网球
func player(name string,game chan int)  {
	defer wg.Done()

	for {
		//等待球过来
		ball,ok := <-game
		if !ok{
			//如果通道关闭，赢了
			fmt.Printf("Player %s Won\n",name)
			return
		}

		n := rand.Intn(100) //随机数判断是否丢球
		if n%13 == 0{
			fmt.Printf("Player %s Missed\n",name)

			//关闭通道，表示输了
			close(game)
			return
		}

		//显示击球数
		fmt.Printf("Player %s Hit %d\n",name,ball)
		ball++

		//将球打向对手
		game <-ball
	}
}
