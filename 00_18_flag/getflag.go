package main

/*
コマンドラインから値の受け取る
*/
import (
	"fmt"
	"flag"
)

func main()  {
	var (
		max int
		msg string
		test string
		x bool
		)

	// ポインタ、オプション、初期値、オプション説明
	flag.IntVar(&max,"n",32,"処理")
	flag.StringVar(&msg,"m","","msg")
	flag.StringVar(&test,"t","11","msg")
	flag.BoolVar(&x,"x",false,"")

	flag.Parse()

	fmt.Println(max)
	fmt.Println(msg)
	fmt.Println(test)
	fmt.Println(x)

}