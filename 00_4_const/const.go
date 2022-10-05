package main

import (
	"fmt"
)

const a int = 100
const (
	c = "TEST"
	d = "KKK"
)

// jk自動的に３になる
const (
   i = 3
   j
   k
)

//iota 0から　１を増やす
const (
	it = iota //0
	it1 //1
	it2 //2
	it3 = 999 //★iota 裏で1を増された
	it4 = iota //4
	it5 //5
)

func main() {
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	fmt.Println(it)
	fmt.Println(it1)
	fmt.Println(it2)
	fmt.Println(it3)
	fmt.Println(it4)
	fmt.Println(it5)
	
}