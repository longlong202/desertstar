/*
如何求解迷宫？
给定一个大小为N*N的迷宫，一只老鼠需要从迷宫的左上角（对应矩阵的[0][0]），走到迷宫的右下角(对应矩阵的[N-1][N-1])，老鼠只能向两个方向移动；向右或向下。在
迷宫中，0表示没有路，1表示有路。请给出算法找出一条合理路径。

主要思路：回溯法
当碰到死胡同的时候，回溯到前一步，然后从前一步出发继续寻找可达的路径。
1、申请一个结果矩阵来标记移动的路径；
2、if到达了目的地；
3、打印解决方案矩阵l
else
1、在结果矩阵中标记当前为1（1表示可以移动的路径）
2、向右前进一步，然后递归地检查，走完这一步后，判断是否存在到终点的可达路线；
3、如果步骤2中的移动方法导致没有通往终点的路径，那么选择向下移动一步，然后检查使用这种移动方法后，是否存在到终点的可达的路线；
4、如果上面的移动方法都会导致没有可达的路径，那么标记当前单元格在结果矩阵中为0，返回false，并回溯到第一步。
 */
package main

import (
	"fmt"
)

var N=4

//打印从起点到终点的路线
func printSolutuon(sol [][]int)  {
	for i:=0;i<4;i++{
		for j:=0;j<N;j++{
			fmt.Print(sol[i][j]," ")
		}
		fmt.Println()
	}
}

//判断x和y是否是一个合理的单元
func isSafe(maze [][]int,x,y int) bool  {
	return (x>=0&&x<N&&y>=0&&y<N&&maze[x][y] ==1)
}

//maze表示迷宫，x,y表示起点，sol存储结果
func getPath(maze [][]int,x,y int,sol [][]int) bool {
	//到达目的地
	if  x==N-1 && y==N-1{
		sol[x][y] = 1
		return true
	}

	//判断maze[x][y]是否是一个可通行的单元
	if isSafe(maze,x,y) {
		//标记当前单元为1
		sol[x][y] =1
		//向右走一步
		if getPath(maze,x,y+1,sol) {
			return true
		}
		//向下走一步
		if getPath(maze,x+1,y,sol) {
			return true
		}

		//标记当前单元为0用来表示这条路径不通,然后回溯
		sol[x][y] =0

	}

	return false
}

func main()  {
	maze := [][]int{
		{1,0,0,0},
		{1,1,0,1},
		{1,0,0,0},
		{1,1,1,1},
	}

	sol := [][]int{
		{0,0,0,0},
		{0,0,0,0},
		{0,0,0,0},
		{0,0,0,0},
	}

	if getPath(maze,0,0,sol){
		printSolutuon(sol)
	}else{
		fmt.Println("不存在可达路径")
	}
}
