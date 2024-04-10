package main

import (
	"fmt"
	"go_algorythm/5Link/Hash/HashTableArray"
)

func main1() {
	fmt.Println(HashTableArray.MySHA("abcd1", 100))
	fmt.Println(HashTableArray.MySHA("abcd", 100))
	fmt.Println(HashTableArray.MySHA("abcd2", 100))
}

// 49
// 96
// 50
func main2() {
	fmt.Println(HashTableArray.MySHA256("abcd1", 100))
	fmt.Println(HashTableArray.MySHA256("abcd", 100))
	fmt.Println(HashTableArray.MySHA256("abcd2", 100))
}
func main() {
	mytable, _ := HashTableArray.NewHashTable(100, HashTableArray.MySHA)
	mytable.Insert("abcd3")
	mytable.Insert("abcd1")
	mytable.Insert("abcd2")
	pos := mytable.Find("abcd1")
	fmt.Println(mytable.GetValue(pos))
	pos = mytable.Find("abcd2")
	fmt.Println(mytable.GetValue(pos))

}
