package main

import (
	"fmt"
)

//队列
func AIoutQueue(AIdatax [M][N]int, i int, j int) bool {
	myq := NewQueue()
	myq.Enqueue(&pos{i, j})

	for {
		nowpos := myq.Dequeue()
		if nowpos == nil {
			break
		}
		if nowpos.x == M-1 && nowpos.y == N-1 {
			i = nowpos.x
			j = nowpos.y
			AIdatax[i][j] = 3 //避免回头路
			cangoout = true
			AIdata = AIdatax
			fmt.Println("迷宫可以走出来")
			break
		} else {
			i = nowpos.x
			j = nowpos.y
			//AIdatax[i][j]=3//避免回头路
			if cangoout {
				AIdatax[i][j] = 0 //走不通设置为0
			} else {
				AIdatax[i][j] = 3 //代表已经走过了，3表示不能回头走，大于2
			}
			if j+1 <= N-1 && AIdatax[i][j+1] < 2 && cangoout != true { //右边
				//AIout(AIdatax,i,j+1) //递归一次
				myq.Enqueue(&pos{i, j + 1})
			}
			if i+1 <= M-1 && AIdatax[i+1][j] < 2 && cangoout != true { //下边
				//AIout(AIdatax,i+1,j) //递归一次
				myq.Enqueue(&pos{i + 1, j})
			}
			if j-1 >= 0 && AIdatax[i][j-1] < 2 && cangoout != true { //左边
				//AIout(AIdatax,i,j-1) //递归一次
				myq.Enqueue(&pos{i, j - 1})
			}
			if i-1 >= 0 && AIdatax[i-1][j] < 2 && cangoout != true { //上边
				//AIout(AIdatax,i-1,j) //递归一次
				myq.Enqueue(&pos{i - 1, j})
			}

		}

	}
	return cangoout
}

//递归--转化为栈
func AIoutStack(AIdatax [M][N]int, i int, j int) bool {

	mystack := NewStack()
	mystack.Push(&pos{i, j})

	for !mystack.IsEmpty() {
		nowpos, err := mystack.Pop()
		if err != nil {
			break //数据取不出来，跳出循环
		}
		i = nowpos.x
		j = nowpos.y
		AIdatax[i][j] = 3 //避免回头路

		if nowpos.x == M-1 && nowpos.y == N-1 {

			cangoout = true
			AIdata = AIdatax
			fmt.Println("迷宫可以走出来")
			break
		} else {

			if j+1 <= N-1 && AIdatax[i][j+1] < 2 && cangoout != true {
				//AIout(AIdatax,i,j+1) //递归一次
				mystack.Push(&pos{i, j + 1})
			}
			if i+1 <= M-1 && AIdatax[i+1][j] < 2 && cangoout != true {
				//AIout(AIdatax,i+1,j) //递归一次
				mystack.Push(&pos{i + 1, j})
			}
			if j-1 >= 0 && AIdatax[i][j-1] < 2 && cangoout != true {
				//AIout(AIdatax,i,j-1) //递归一次
				mystack.Push(&pos{i, j - 1})
			}
			if i-1 >= 0 && AIdatax[i-1][j] < 2 && cangoout != true {
				//AIout(AIdatax,i-1,j) //递归一次
				mystack.Push(&pos{i - 1, j})
			}

		}

	}

	return cangoout

}

//解决递归走出来
func AIout(AIdatax [M][N]int, i int, j int) bool {
	AIdatax[i][j] = 3 //避免回头路
	//AIdatax[i][j]=3//避免回头路
	//show(AIdatax)
	if i == M-1 && j == N-1 {
		cangoout = true
		AIdata = AIdatax
		fmt.Println("迷宫可以走出来")
	} else {
		if j+1 <= N-1 && AIdatax[i][j+1] < 2 && cangoout != true {
			AIout(AIdatax, i, j+1) //递归一次
		}
		if i+1 <= M-1 && AIdatax[i+1][j] < 2 && cangoout != true {
			AIout(AIdatax, i+1, j) //递归一次
		}
		if j-1 >= 0 && AIdatax[i][j-1] < 2 && cangoout != true {
			AIout(AIdatax, i, j-1) //递归一次
		}
		if i-1 >= 0 && AIdatax[i-1][j] < 2 && cangoout != true {
			AIout(AIdatax, i-1, j) //递归一次
		}

	}
	return cangoout
}
func AImoveOut() {
	AIdata[ipos][jpos] = 1
	run("")
	//ipos+=1
	//ipos=a[9][9]

	for ipos != 9 || jpos != 9 {

		if ipos-1 >= 0 && AIdata[ipos-1][jpos] == 3 {
			AIdata[ipos-1][jpos] = 0
			run("w")
		}
		if ipos+1 <= 9 && AIdata[ipos+1][jpos] == 3 {
			AIdata[ipos+1][jpos] = 0
			run("s")
		}
		if jpos+1 <= 9 && AIdata[ipos][jpos+1] == 3 {
			AIdata[ipos][jpos+1] = 0
			run("d")
		}
		if jpos-1 >= 0 && AIdata[ipos][jpos-1] == 3 {
			AIdata[ipos][jpos-1] = 0
			run("a")
		}

		fmt.Println(ipos, jpos)
	}
}
