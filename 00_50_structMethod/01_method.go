package main

import "fmt"

// Point 構造体アリアス
type Point struct {
	x, y int
}

// Sum ＳＵＭメソッド作成　ポインタ
func (p *Point) Sum(x int, y int) int {
	return x + y
}

// 関数型
func (p Point) xySum() int {
	return p.x + p.y
}

// 型のメソッド
type myInt int

func (i myInt) intSum(j myInt) myInt {
	return i + j
}

// 配列
type array [3]int

func (ar array) first() int {
	return ar[0]
}

func (ar array) last() int {
	return ar[2]
}

// 配列
type Strings []string

func (s Strings) join(d string) string {
	sumStr := ""
	for _, v := range s {
		if sumStr != "" {
			sumStr += d
		}
		sumStr += v
	}
	return sumStr
}

// 関数としてのメソッド
type Student struct {
	id   string
	name string
}

func (s *Student) getIdAndName() string {
	return s.id + "/" + s.name
}

func main() {
	fmt.Println(123)
	// ポインタから呼び出す
	p := &Point{1, 2}
	sum := p.Sum(1, 2)
	fmt.Println(sum)

	// 関数から呼び出す
	p1 := Point{5, 6}
	sum1 := p1.xySum()
	fmt.Println(sum1)

	// 型エリアすのメソッド呼び出す
	intSum := myInt(4).intSum(5)
	fmt.Println(intSum)

	// 配列の値 数字
	first := array{1, 2, 3}.first()
	fmt.Println(first)
	last := array{9, 8, 7}.last()
	fmt.Println(last)

	// Ｓｔｒｉｎｇスライス
	strings := Strings{"A", "4", "B", "C"}
	var ret string
	ret = strings.join("--")
	fmt.Printf("ret=%s\n", ret)

	// 関数としてとメソッド
	f := (*Student).getIdAndName
	// 関数として呼び出す
	str := f(&Student{"1", "2"})
	fmt.Printf("StudentName = %s", str)
}
