package main

import (
	"fmt"
	"log"
	"os"
)

// fmt参照
// https://qiita.com/rock619/items/14eb2b32f189514b5c3c

func main() {
	f, err := os.Create("foo.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(f, "%05d\n", 123)
	defer f.Close()
	lf, _ := os.Create("log.txt")
	defer f.Close()

	log.SetOutput(lf)
	log.Println("log")
}
