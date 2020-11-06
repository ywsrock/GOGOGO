package main

import (
    "fmt"
)

type Student struct {
    id int
    name string
    age int
}

type Class struct {
    id int
    students []Student
}

func main() {
    // スライスサクセイ　値の設定
    var students1 = make([]Student, 1)
    students1[0].id=1
    students1[0].name="test"
    students1[0].age=18

   //スライスの作成 スライスマージ
   var students []Student
   students = append(students,students1...)

	//構造体の作成、初期化
   var st2 Student = Student{1,"st2",18}
   var st3 Student = Student{id:2,name:"st3",age:12}

   //スライスに格納
   students=append(students,st2,st3)
   
    //構造体に値の設定
    var class1 Class
    // クラスＩＤ
    class1.id = 1
    //学生の設定
    class1.students=students
    fmt.Println("classID=",class1.id)

    //学生一覧 
    for _,v := range class1.students{
        fmt.Println(v.id)
        fmt.Println(v.name)
        fmt.Println(v.age)
        fmt.Println("-----")
    }
}
