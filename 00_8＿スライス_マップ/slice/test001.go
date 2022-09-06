package main

import "fmt"

func raise(s []int) {
	fmt.Println("Len:",len(s))
	fmt.Println("Cap:",cap(s))
	s = append(s,0)

	for i := range s{
		s[i] ++
	}
}


func main () {
	s1 := []int{1,2}
	s2 := s1
	s2 = append(s2,3)
	raise(s1)
	raise(s2)

	fmt.Println(s1,s2)



}
