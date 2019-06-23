//使用有缓冲的通道和固定数目的goruntine来处理事情
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4//要使用的gorountine的数量
	taskLoad = 10 //要处理的工作的数量
)

//使用wg来等待程序的完成
var wg sync.WaitGroup

//优先执行初始化包
func init()  {
	rand.Seed(time.Now().Unix())
}

func main() {
	//创建一个有缓冲的通道来工作
	tasks := make(chan string,taskLoad)

	//启动goroutine来处理工作
	wg.Add(numberGoroutines)

	for gr := 1;gr <= numberGoroutines;gr++ {
		go worker(tasks,gr)
	}

	//增加一组要完成的工作
	for post := 1; post <= taskLoad;post++{
		tasks <- fmt.Sprintf("Task : %d",post) //将任务传入通道
	}

	//当所有工作都处理完再关闭通道
	close(tasks)

	//等待所有工作完成
	wg.Wait()
}

//处理有缓冲通道传入的工作
func worker(tasks chan string,worker int)  {

	//通知函数已经返回
	defer wg.Done()

	for {
		task,ok := <-tasks
		if !ok {
			//如果通道接收传回的ok状态是false，意味着通道已经没有值，通过ok标识判断通道是否还有值
			fmt.Printf("Worker : %d : Shutting Down\n",worker)
			return
		}

		//开始工作
		fmt.Printf("Worker : %d : Startted %s\n",worker,task)

		//随机等待一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		//完成工作
		fmt.Printf("Worker : %d : Completed %s \n",worker,task)
	}
}

