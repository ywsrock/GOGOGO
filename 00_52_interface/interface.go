package main

import (
	"fmt"
	"reflect"
)

type popAction interface {
	getEmail() string
}

type Teacher struct {
	name string
}

type Student struct {
	name string
	age  int
}

func (t Teacher) getEmail() string {
	return t.name + "@teacher.com"
}

func (t Student) getEmail() string {
	return t.name + "@student.com"
}

func sendEmail(p popAction) {
	fmt.Printf("EMAIL=%s \n", p.getEmail())
}

func checkType(p popAction) string {
	// 型にチェック
	switch p.(type) {
	case Teacher:
		return "Teacher"
	default:
		return "Student"
	}
}

func main() {
	s := Student{"su", 1}
	t := Teacher{"Tea"}
	sendEmail(s)
	sendEmail(t)

	var i popAction
	// 型のアサーション
	i = s
	fmt.Println(i)
	_, ok := i.(Student)
	_, isT := i.(Teacher)
	fmt.Println(ok)
	fmt.Println(isT)
	//  型switch
	fmt.Println(checkType(s))
	fmt.Println(checkType(t))

	//reflect
	// var f popAction
	var f interface{}
	// f = s
	v := reflect.ValueOf(f)
	fmt.Println("-----------------------------")
	fmt.Println(v)
	fmt.Println(v.IsValid())
}
