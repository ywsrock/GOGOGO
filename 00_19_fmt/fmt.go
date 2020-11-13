package main

import (
	"fmt"
	"os"
	"log"
)

func main()  {
	f,err := os.Create("foo.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(f,"%05d\n",123)

	lf ,err := os.Create("log.txt")
	log.SetOutput(lf)
	log.Println("log")
}
