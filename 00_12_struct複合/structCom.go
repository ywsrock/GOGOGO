package main

import "fmt"

//Class 構造体
type Class struct {
	id   int
	name string
}

//School 構造体
type School struct {
	id    int
	name  string
	class []Class //複数Classを含む
}

//Student 構造体
type Student struct {
	Class       //fieldメイ指定なしの場合、直接利用できる(★field名を重複しないように注意)
	studentName string
	studentID   int
}

func main() {
	fmt.Println("-----構造複合-----")
	// 構造初期化
	// class := make([]Class, 1)
	var classes []Class
	classes = append(classes, Class{1, "class1"})
	classes = append(classes, Class{2, "class2"})
	classes = append(classes, Class{3, "class3"})
	classes = append(classes, Class{4, "class5"})

	school := School{id: 1, name: "school", class: classes}
	//school出力
	fmt.Println(school)
	//class2の値出
	fmt.Println(school.class[1].name)

	//フィールド名指定しない場合、直接アクセスできます
	class := Class{1, "Class-Student"}
	student := Student{studentID: 1, studentName: "student1", Class: class}
	//クラス名をアクセス
	fmt.Println(student.name)
}
