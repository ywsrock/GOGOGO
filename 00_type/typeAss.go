package  main

import  "fmt"

//多い方
type iPerson interface {
	m1()
	m2()
}

//少ない方
type IPerson2 interface {
	m1 ()
	//m2 ()
}

type P1 struct {
	name string
	age string
}

	type P2 struct {
	name string
	age string
}

func (p P1) m1()  {
	fmt.Println(p.name)
}
func (p P1) m2()  {
	fmt.Println(p.name)
}

func (p P2) m1()  {
	fmt.Println(p.name)
}
func (p P2) m2()  {
	fmt.Println(p.name)
}

func Print(in iPerson)  {
	p,ok := in.(IPerson2)
	fmt.Println(ok)
	fmt.Printf("%#v\n",p)
	in.m1()
}

func main(){
	p1 := P1{name:"yws",age:"123"}
	p2 := P2{"aaa","333"}

	Print(p1)
	Print(p2)



}