package main

import (
	"./bar"
	"fmt"
)

func main() {
	str := bar.DoTest()
    fmt.Println(str)
}