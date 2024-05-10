package main

import "fmt"

//"wsad "//w 上，s下，A，左，d右边
func run(direct string) {

	if direct == "w" {
		if ipos-1 >= 0 && data[ipos-1][jpos] < 2 {
			//交换数据
			data[ipos][jpos], data[ipos-1][jpos] = data[ipos-1][jpos], data[ipos][jpos]
			ipos -= 1
		}
	} else if direct == "s" {
		if ipos+1 <= M-1 && data[ipos+1][jpos] < 2 {
			data[ipos][jpos], data[ipos+1][jpos] = data[ipos+1][jpos], data[ipos][jpos]
			ipos += 1
		}

	} else if direct == "a" {
		if jpos-1 >= 0 && data[ipos][jpos-1] < 2 {
			data[ipos][jpos], data[ipos][jpos-1] = data[ipos][jpos-1], data[ipos][jpos]
			jpos -= 1
		}

	} else if direct == "d" {
		if jpos+1 <= N-1 && data[ipos][jpos+1] < 2 {
			data[ipos][jpos], data[ipos][jpos+1] = data[ipos][jpos+1], data[ipos][jpos]
			jpos += 1
		}
	} else {

	}
	fmt.Println("-----------------------data")
	//fmt.Println(direct,ipos,jpos)
	show(data)
}
