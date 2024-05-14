package main

import "math"

//A星算法地图上的点结构
type AstarPoint struct {
	Point
	father *AstarPoint
	gVal   int //g(n)表示从初始结点到任意结点n的代价
	hVal   int //h(n)表示从结点n到目标点的启发式评估代价
	fVal   int //f(n)综合评估最小的结点n，其中f(n) = g(n) + h(n)
}

func NewAstarPoint(p *Point, father *AstarPoint, end *AstarPoint) (ap *AstarPoint) {
	ap = &AstarPoint{*p, father, 0, 0, 0} //初始化
	if end != nil {
		ap.CalcFval(end) //创建的时候就计算节点的评估
	}
	return ap
}

// x y  -
//
//g(n)表示从初始结点到任意结点n的代价
func (this *AstarPoint) CalcGval() int {
	if this.father != nil {
		deltaX := math.Abs(float64(this.father.x - this.x))
		deltaY := math.Abs(float64(this.father.y - this.y)) //从父亲节点走过来需要的代价
		if deltaX == 1 && deltaY == 0 {
			//移动一步
			this.gVal = this.father.gVal + 10
		} else if deltaX == 0 && deltaY == 1 {
			this.gVal = this.father.gVal + 10
		} else if deltaX == 1 && deltaY == 1 {
			this.gVal = this.father.gVal + 14
		} else {
			panic("error")
		}

	}

	return this.gVal //返回
}

//计算当前节点与目标节点的差距
func (this *AstarPoint) CalcHval(end *AstarPoint) int {
	this.hVal = int(math.Abs(float64(end.x-this.x)) + math.Abs(float64(end.y-this.y)))
	return this.hVal
}
func (this *AstarPoint) CalcFval(end *AstarPoint) int {
	this.fVal = this.CalcGval() + this.CalcHval(end)
	return this.fVal
}
