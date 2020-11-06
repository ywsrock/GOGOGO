package main

//sub package のメソッド呼び出す
import (
	"fmt"
	"./subpackage"
)

func main()  {
	fmt.Println("this first")	
	fmt.Println(subpackage.Sub001())

}