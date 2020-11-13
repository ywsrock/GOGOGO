package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func writeFile(fname string) {
	ioutil.WriteFile(fname, []byte("1234"), 666)
}
func readFile(fname string) {
	bs, _ := ioutil.ReadFile(fname)
	fmt.Println(string(bs))
}

func readALL(fname string) {
	f, _ := os.Open(fname)
	bs, _ := ioutil.ReadAll(f)
	fmt.Println(string(bs))
}

func main() {
	writeFile("123.txt")
	readFile("123.txt")
	readALL("123.txt")
}
