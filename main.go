package main

import "fmt"

// 食べるためのインターフェース
type Eater interface {
	PutIn()   // 口に入れる
	Chew()    // 噛む
	Swallow() // 飲み込む
}

// 人間の構造体
type Human struct {
	Height int // 身長
}

// カメの構造体
type Turtle struct {
	Kind string // 種類
}

// 人間用のインターフェースの実装
func (h Human) PutIn() {
	fmt.Println("道具を使って丁寧に口に運ぶ")
}
func (h Human) Chew() {
	fmt.Println("歯でしっかり噛む")
}

func (h Human) Swallow() {
	fmt.Println("よく噛んだら飲み込む")
}

// カメ用のインターフェースの実装
func (h Turtle) PutIn() {
	fmt.Println("獲物を見つけたら首をすばやく伸ばして噛む")
}
func (h Turtle) Chew() {
	fmt.Println("クチバシで噛み砕く")
}
func (h Turtle) Swallow() {
	fmt.Println("小さく砕いたら飲み込む")
}

// インターフェースが引数になる、食べるメソッド
func EatAll(e Eater) {
	e.PutIn() // インターフェースから呼び出す
	e.Chew()
	e.Swallow()
}

func main() {
	var man Human = Human{Height: 300}         // 人間用の構造体を作成
	var cheloniaMydas = Turtle{Kind: "アオウミガメ"} // カメ用の構造体を作成
	var eat Eater                              // インターフェースEater型の変数を用意
	fmt.Println("＜人間が食べる＞")
	eat = man   // Human型がインターフェースであるEater型に変換される
	EatAll(eat) // インターフェースを引数に関数を呼ぶ
	fmt.Println("＜カメが食べる＞")
	eat = cheloniaMydas // Turtle型がインターフェースであるEater型に変換される
	EatAll(eat)
}
