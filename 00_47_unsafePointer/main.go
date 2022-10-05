package main

import (
	"fmt"
	"reflect"
	"unsafe"
)
func at(s reflect.SliceHeader, i int) unsafe.Pointer {

	fmt.Printf("--- %v",uintptr(0))
	fmt.Printf("---??? %v",unsafe.Sizeof(0))

	// 先頭ポインタ + インデックス * int型のサイズ
	return unsafe.Pointer(s.Data + uintptr(i)*unsafe.Sizeof(int(0)))
}

func main() {
	a := [...]int{10, 20, 30}

	t:= unsafe.Pointer(&a[0])
	fmt.Printf("unsafe.P %v",t)

	uintP := uintptr(unsafe.Pointer(&a[0]))
	fmt.Printf("unsafe.ini %v",uintP)

	s := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&a[0])), Len: 2, Cap: len(a)}
	*(*int)(at(s, 0)) = 100 // unsafe.Pointerを*intに変換して代入している
	fmt.Println(a)
}
