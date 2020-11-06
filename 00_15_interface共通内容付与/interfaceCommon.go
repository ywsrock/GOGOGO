package main

import "fmt"

//say interface (共通の動作を規定する)
type mouth interface {
	say() string
}

//人間構造
type human struct {
	name string
}

//動物
type animal struct {
	name string
}

//human say　の実装
func (h *human) say() string {
	return "人間の言葉"
}

//animal say の実装
func (a *animal) say() string {
	return "動物語"
}

//human と　animal共通のメソッド付与
func useMouth(mt mouth) string {
	return mt.say()
}

func main() {
	man := &human{"superMan"}
	saylang := useMouth(man)
	fmt.Println(saylang)

	dog := &animal{"superDog"}
	saylang = useMouth(dog)
	fmt.Println(saylang)
}
