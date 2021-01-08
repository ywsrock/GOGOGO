package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Stringへに変換
	i := strconv.FormatInt(123, 2)
	fmt.Println(i)

	j := strconv.FormatInt(123, 16)
	fmt.Println(j)

	f := strconv.FormatInt(123, 10)
	fmt.Println(f)

	//Stringからの変換
	a, _ := strconv.ParseBool("true")
	b, _ := strconv.ParseInt("123", 10, 0)
	fmt.Println(a)
	fmt.Println(b)
}
