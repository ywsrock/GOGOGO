package main

import "fmt"

func main() {
	s1 := []int{1, 2}
	fmt.Println("s1", len(s1), cap(s1))

	s2 := s1
	fmt.Println("s2", len(s2), cap(s2))

	// 容量拡張　現在容量の2倍になる
	s2 = append(s2, 3)
	fmt.Println("s3", len(s2), cap(s2))

	// s1容量拡張のため、メソッドのsとs1は違います。（内部は参照した配列はコピー発生）
	SliceRise(s1)
	// 容量が4なので、s2はの容量拡張なし、同じ
	SliceRise(s2)

	fmt.Println(s1, s2)

}

func SliceRise(s []int) {
	s = append(s, 0)
	for i := range s {
		s[i]++
	}
}
