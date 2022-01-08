package main

import "fmt"

func main() {
	s1 := []int{1, 2}
	fmt.Println("s1",len(s1),cap(s1))

	s2 := s1
	fmt.Println("s2",len(s2),cap(s2))



	s2 = append(s2, 3)
	fmt.Println("s3",len(s2),cap(s2))



	SliceRise(s1)
	SliceRise(s2)

	fmt.Println(s1,s2)

}

func SliceRise(s []int) {
	s = append(s, 0)

	for i := range s {
		s[i]++
	}
}
