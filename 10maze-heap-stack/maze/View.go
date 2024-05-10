package main

import "fmt"

func show(arr [M][N]int) {
	//fmt.Println("-----------------------------------------")
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%4d", arr[i][j])
		}
		fmt.Println("")
	}

}
