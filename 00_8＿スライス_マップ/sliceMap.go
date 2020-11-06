package main

import "fmt"

// 複数パラメータ渡す
func sum(s ...int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("参照型")

	//スライスの定義
	var s []int
	//配列の定義
	var s1 [3]int
	//配列かスライス
	s = s1[0 : len(s1)-1]

	fmt.Println(s)
	fmt.Println(s1)

	//スライス初期値設定
	var s2 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(s2)

	// makeからスライスの作成
	s3 := make([]int, 3)
	fmt.Println(s3)

	//★スライスの罠 一部コピーや全部コピーの場合、参照先変わると、スライスの値も変わります。
	ar1 := [5]int{1, 3, 4, 5, 6}
	sl1 := ar1[0:3]
	fmt.Println("スライス作成時:", sl1)
	// 参照先の値を変更すると、スライスの値の変わる
	ar1[1] = 2
	fmt.Println("参照先変更するにもスライスにも", sl1)

	//スライスのappend関数を利用して、スライスのコピーを発生させる
	ar2 := [5]int{1, 2, 3, 4, 5}
	//★★全部コピーしないと、上記と同じになる
	sl2 := ar2[:]
	fmt.Println(sl2)
	sl2 = append(sl2, 99)
	fmt.Println(sl2)
	//参照先変更する
	ar2[1] = 5
	fmt.Println(sl2)

	//複数パラメータの渡す
	ret := sum(1, 2, 3)
	fmt.Println(ret)
	ret = sum(1)
	fmt.Println(ret)

	//...により、スライス引数
	sl3 := []int{1, 3, 4, 5, 6}
	ret = sum(sl3...)
	fmt.Println(ret)

	//マップ
	//新規作成
	var m map[int]string = map[int]string{1: "A"}
	fmt.Println("マップ:", m)

	//リテラルでmap作成
	m1 := map[int]string{1: "A", 2: "B"}
	fmt.Println("マップlen-->", len(m1))
	fmt.Println(m1)
	//指定キー削除処理
	delete(m1, 2)
	fmt.Println("マップdelete-->", m1)

	//make　からマップ作成
	m2 := make(map[int]string)
	m2[1] = "string"
	fmt.Println(m2)

	//存在しないキーをアクセスしても、エラーにならない
	fmt.Println("存在しないキー:", m2[2])
	// 存在しないキーの判断するため、下記の書き方
	mv, isOk := m2[2]
	if isOk {
		fmt.Println(mv)
	} else {
		fmt.Println("キー存在しません......")
	}

	//マップloop
	for k, v := range m2 {
		fmt.Println("k-->", k, "v-->", v)
	}
}
