package main

import "fmt"

// 二つInterfaceに関して、型アサーションの検証
// .IF IPerson {m1(),m2()}
// .IF Iperson2 {m1()}

// IPerson から　IPerson2にassertionはOK(Iperson2のm1メソッドがふくまれたので、OK)
// IPerson2 から　IPersonにassertionはNG(m2()がないので、NG）

// 多い方
type IPerson interface {
	m1()
	m2()
}

// 少ない方
type IPerson2 interface {
	m1()
	// m2 ()
}

type P1 struct {
	name string
	age  string
}

type P2 struct {
	name string
	age  string
}

func (p P1) m1() {
	fmt.Println(p.name)
}
func (p P1) m2() {
	fmt.Println(p.name)
}

func (p P2) m1() {
	fmt.Println("p2.m1()", p.name)
}

func (p P2) m2() {
	fmt.Println("p2.m2()", p.name)
}

func Print(in IPerson) {
	// type assertion
	p, ok := in.(IPerson2)
	fmt.Println(ok)
	fmt.Printf("%#v\n", p)
	// in.m1()
	p.m1()
	// p.m2() エラー
}

func main() {
	p1 := P1{name: "yws", age: "123"}
	p2 := P2{"aaa", "333"}
	Print(p1)
	Print(p2)
}
