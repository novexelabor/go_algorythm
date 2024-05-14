package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PointAsKey(x, y int) (key string) {
	key = strconv.Itoa(x) + "," + strconv.Itoa(y) //坐标转化为字符串
	return key
}

type Map struct {
	points [][]Point         //地图，保存矩阵
	blocks map[string]*Point //字符串对应每个节点,标记障碍
	maxX   int               //最大的X坐标
	maxY   int               //最大的Y坐标
}

func NewMap(charMap []string) (m Map) {
	m.points = make([][]Point, len(charMap))           //开辟内存存储二维数组
	m.blocks = make(map[string]*Point, len(charMap)*2) //两倍边长
	for x, row := range charMap {
		cols := strings.Split(row, " ")        //基于空格切割
		m.points[x] = make([]Point, len(cols)) //二维数组每个元素开辟内存
		for y, view := range cols {
			m.points[x][y] = Point{x, y, view}
			if view == "X" {
				//标记障碍
				m.blocks[PointAsKey(x, y)] = &m.points[x][y]
			}
		}
	}
	m.maxX = len(m.points) //取得边界
	m.maxY = len(m.points[0])
	return m
}

//抓取相邻节点,返回一个集合
//   x-1,y+1  x,y+1  x+1,y+1
//   x-1,y    x,y    x+1,y
//   x-1,y-1  x,y-1 x+1,y-1
func (this *Map) GetAdjaoentPoint(curPoint *Point) (Adjaoent []*Point) {
	if x, y := curPoint.x, curPoint.y-1; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		Adjaoent = append(Adjaoent, &this.points[x][y])
	}
	if x, y := curPoint.x+1, curPoint.y-1; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		Adjaoent = append(Adjaoent, &this.points[x][y])
	}
	if x, y := curPoint.x+1, curPoint.y; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		Adjaoent = append(Adjaoent, &this.points[x][y])
	}
	if x, y := curPoint.x+1, curPoint.y+1; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		Adjaoent = append(Adjaoent, &this.points[x][y])
	}
	if x, y := curPoint.x, curPoint.y+1; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		Adjaoent = append(Adjaoent, &this.points[x][y])
	}
	if x, y := curPoint.x-1, curPoint.y+1; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		Adjaoent = append(Adjaoent, &this.points[x][y])
	}
	if x, y := curPoint.x-1, curPoint.y; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		Adjaoent = append(Adjaoent, &this.points[x][y])
	}
	if x, y := curPoint.x-1, curPoint.y-1; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		Adjaoent = append(Adjaoent, &this.points[x][y])
	}
	return Adjaoent
}

//打印地图，按照寻找好的通路
func (this *Map) printMap(path *SearchRoad) {
	fmt.Println("地图的边界", this.maxX, this.maxY)
	for x := 0; x < this.maxX; x++ {
		for y := 0; y < this.maxY; y++ {
			if path != nil {
				if x == path.start.x && y == path.start.y {
					fmt.Printf("%2s", "S") //S代表开始
					goto NEXT
				}
				if x == path.end.x && y == path.end.y {
					fmt.Printf("%2s", "E") //E代表结束
					goto NEXT
				}
				for i := 0; i < len(path.TheRoad); i++ { //循环找路
					if path.TheRoad[i].x == x && path.TheRoad[i].y == y {
						fmt.Printf("%2s", "*") //*代表走过
						goto NEXT
					}
				}
			}
			fmt.Printf("%2s", this.points[x][y].view)

		NEXT:
		}
		fmt.Println()
	}

}
