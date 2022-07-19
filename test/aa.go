package main

import "fmt"

type A struct {
	k, val string
}

func main() {

	v := &A{}
	v.k = "1"
	v.val = "2"
	if true {
		fmt.Println(v.val)
	}
}
