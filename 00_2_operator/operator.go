package main

import (
	"fmt"
)

func main() {
	fmt.Println("他の言語と一緒です")
	// 算術 + - * / % << >> | & ^
	// 比較 > < = >= <= !=
	// 論理 && ||
	// 優先順位  * / % + -  ==
	a := 1 + 1
	fmt.Println(a)

	var b int
	b = 2
	b += 1
	fmt.Println(b)

	fmt.Println(a > b)
	fmt.Println(a != b)

    a1 := 1 != 1 && 2 != 2
	fmt.Println(a1)

}
