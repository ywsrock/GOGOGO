package main

import ("fmt")

type Point struct{
	x int
	y int
}

// メソッド定義 Point 型
func (p Point) getPoint(x int, y int)Point  {
	p.x = x
	p.y = y
	return p
}

// 戻り値ない、ポインタ型
func (p *Point) pPoint(x int,y int)  {
	p.x = x
	p.y = y
}

func main()  {

	//★呼び出す側には、ポインタ型でも、値型でも、関数呼び出すことができます
	p := Point{}
	retP := p.getPoint(2,3)
	fmt.Printf("Px=%d,Py=%d \n",retP.x,retP.y)

    // 値型でポインタ型のレシーバを呼び出す
	p.pPoint(9,10)
	fmt.Printf("px=%d,py=%d",p.x,p.y)

	// 同じポインタ型で、メソッドを呼び出す、戻り値なし
	p1 := &Point{}
	p1.pPoint(3,5)
	fmt.Printf("Px=%d,Py=%d \n",p1.x,p1.y)

	// ポインタ型で、値型のレシーバを呼び出す
	retp2 := p1.getPoint(123,99)	
	fmt.Printf("Px=%d,Py=%d",retp2.x,retp2.y)
}