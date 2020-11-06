package main

import (
	"strconv"
	"fmt"
)

func main()  {
	i := strconv.FormatInt(123,2)
	fmt.Println(i)
	j := strconv.FormatInt(123,16)
	fmt.Println(j)
	f := strconv.FormatInt(123,10)
	fmt.Println(f)
}