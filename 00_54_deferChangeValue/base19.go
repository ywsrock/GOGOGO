package main

import "fmt"

type Person struct {
	age int
}

func main() {
	person := &Person{28}

	// 1.変更後値出力
	defer fmt.Println(person.age)

	// 2. 変更後値
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3.変更前
	defer func() {
		fmt.Println(person.age)
	}()

	// person = &Person{29}
	person.age = 29
}
