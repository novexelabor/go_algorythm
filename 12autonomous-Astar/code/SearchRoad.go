package main

import "container/heap"

type SearchRoad struct {
	theMap  *Map                   //地图
	start   AstarPoint             //开始
	end     AstarPoint             //结束
	closeLi map[string]*AstarPoint //关闭，不通的路
	openLi  OpenList               //通路
	openSet map[string]*AstarPoint //去掉重复
	TheRoad []*AstarPoint          //通道
}

func NewSearchRoad(startx, starty, endx, endy int, m *Map) *SearchRoad {
	sr := &SearchRoad{}
	sr.theMap = m //设定地图
	//开始节点
	sr.start = *NewAstarPoint(&Point{startx, starty, "S"}, nil, nil)
	//结束节点
	sr.end = *NewAstarPoint(&Point{endx, endy, "E"}, nil, nil)
	sr.TheRoad = make([]*AstarPoint, 0)                      //路，开辟内存
	sr.openSet = make(map[string]*AstarPoint, m.maxX+m.maxY) //开放集合
	sr.closeLi = make(map[string]*AstarPoint, m.maxX+m.maxY)
	heap.Init(&sr.openLi)                                      //初始化栈
	heap.Push(&sr.openLi, &sr.start)                           //压入开始节点
	sr.openSet[PointAsKey(sr.start.x, sr.start.y)] = &sr.start //开放集合压入开始节点

	//所有的障碍加入blocks
	for k, v := range m.blocks {
		sr.closeLi[k] = NewAstarPoint(v, nil, nil)
	}
	return sr

}

//A星算法核心
func (this *SearchRoad) FindoutShortestPath() bool {
	//如果开放节点大于0.永远循环下去，
	for len(this.openLi) > 0 {
		//从开方节点中取出放入关闭节点
		x := heap.Pop(&this.openLi)                                 //取出一个节点
		curPoint := x.(*AstarPoint)                                 //取得当前节点
		delete(this.openSet, PointAsKey(curPoint.x, curPoint.y))    //删除开放列表
		this.closeLi[PointAsKey(curPoint.x, curPoint.y)] = curPoint //障碍走过的路，加入关闭列表

		adjacs := this.theMap.GetAdjaoentPoint(&curPoint.Point) //取出所有的邻居节点
		for _, p := range adjacs {
			theAP := NewAstarPoint(p, curPoint, &this.end) //创建A型节点
			//我们找到了结束节点
			if PointAsKey(theAP.x, theAP.y) == PointAsKey(this.end.x, this.end.y) {
				for theAP.father != nil {
					this.TheRoad = append(this.TheRoad, theAP) //加入节点
					theAP.view = "*"                           //标记我们走过
					theAP = theAP.father                       //返回上一个节点
				}
				return true //结束
			}

			_, ok := this.closeLi[PointAsKey(p.x, p.y)] //节点已经存在提前结束本次循环
			if ok {
				continue
			}

			existAP, ok := this.openSet[PointAsKey(p.x, p.y)] //取出开放的节点，如果节点存在，不存在
			if !ok {
				heap.Push(&this.openLi, theAP)                     //节点不存在就压入
				this.openSet[PointAsKey(theAP.x, theAP.y)] = theAP //放入开房列表
			} else {
				//如果节点存在,经过对比取得最短路径
				oldgvar, oldfather := existAP.gVal, existAP.father
				existAP.father = curPoint //当前的父亲节点
				existAP.CalcGval()        //计算最短的值

				//新的节点距离比老节点短
				if existAP.gVal > oldgvar {
					//保存最短的
					existAP.father = oldfather
					existAP.gVal = oldgvar
				}

			}

		}

	}

	return false
}
