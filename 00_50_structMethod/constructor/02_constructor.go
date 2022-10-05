package main

import ("fmt")

//class構造体
type class struct{
	id int
	name string
	count int
}

//class 初期化
func NewClass(id int,name string,count int) *class {
	//new クラスのポインタを返す
	cls := new(class)
	cls.id = id
	cls.name = name
	cls.count = count
	return cls
}


func main()  {
	// 初期化関数を呼び出す
	cls := NewClass(1,"class 001",100)
	// 結果の出力
	fmt.Printf("className=%s",cls.name)
}