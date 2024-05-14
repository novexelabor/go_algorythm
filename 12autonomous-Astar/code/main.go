package main

import "fmt"

func main() {
	presetMap := []string{
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		"X . X X X X X X X X X X X X X X X X X X X X X X X X X",
		". . . . . . . . . . . . . . X . . . . . . . . . . . .",
		". . . . . . . . . . . . . . X . . . . . . . . . . . .",
		". . . . . . . . . . . . . . X . . . . . . . . . . . .",
		". . . . . . . . . . . . . . X . . . . . . . . . . . .",
		". . . . . . . . . . . . . . X . . . . . . . . . . . .",
		". X X X X X X X X X X X X X X . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		"X X X X X X X X X X X X X X X X X X X X X X X X . X X",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		"X . X X X X X X X X X X X X X X X X X X X X X X X X X",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
	}
	m := NewMap(presetMap) //初始化
	m.printMap(nil)
	searchroad := NewSearchRoad(0, 0, 18, 18, &m) //寻路
	if searchroad.FindoutShortestPath() {
		fmt.Println("找到")
		m.printMap(searchroad)
	} else {
		fmt.Println("找不到")
	}

}
