package main

import "fmt"

func isBool(n int) bool {
	if n == 0 {
		return false
	}
	return true

}
func main() {
	ret := isBool(1)
	fmt.Println(ret)
}
