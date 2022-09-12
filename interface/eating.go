// interfaceよくわからないので深堀
// https://qiita.com/rtok/items/46eadbf7b0b7a1b0eb08

package main
import "fmt"

// 食べるためのインターフェイス（人もカメも共通）
type Eater interface {
	PutIn() // 口に入れる
    Chew() // 噛む
    Swallow() // 飲み込む
}

// 人間の構造体
type Human struct {
	Height int // 身長
}

// カメの構造体
type Turtle struct{
    Kind string // 種類
}

// 人間用のインターフェイス実装（人間固有）
// 明示的にEaterインターフェイスに実装するという宣言をしなくても自動的に実装される
func (h Human) PutIn() {
	fmt.Println("道具を使って丁寧に口に運ぶ")
}
func (h Human) Chew(){
    fmt.Println("歯でしっかり噛む")
}
func (h Human) Swallow(){
    fmt.Println("よく噛んだら飲み込む")
}

// カメ用のインターフェイス実装（カメ固有）
// 明示的にEaterインターフェイスに実装するという宣言をしなくても自動的に実装される
func (h Turtle) PutIn(){
    fmt.Println("獲物を見つけたら首をすばやく伸ばして噛む")
}
func (h Turtle) Chew(){
    fmt.Println("クチバシで噛み砕く")
}
func (h Turtle) Swallow(){
    fmt.Println("小さく砕いたら飲み込む")
}

// インターフェイスが引数になる食べるメソッド（人間、カメ共通）
// 食べるのに必要な作業をインターフェイスから人間、カメ意識せずに呼び出せる
func EatAll(e Eater) {
	e.PutIn() // インターフェースから呼び出す
    e.Chew()
    e.Swallow()
}

// インターフェイスが引数になる飲むメソッド（人間、カメ共通）
// 飲むのに必要な作業をインターフェイスから人間、カメ意識せずに呼び出せる
// 共通処理としてインターフェイスを使いまわせる
func DrinkAll(e Eater) {
	e.PutIn() // インターフェースから呼び出す
	e.Swallow()
}

func main(){
	var man Human = Human{Height: 300} // 人間用の構造体を作成
	var cheloniaMydas Turtle = Turtle{Kind: "アオウミガメ"} // カメ用の構造体を作成

	var eat Eater // ★インターフェースEater型の変数を用意

	eat = man // Human型がインターフェースであるEater型に変換される（Human型にはインターフェイスが実装されている）
	fmt.Println("***人間***")
	fmt.Println("***人間の食べる処理***")
	EatAll(eat)
	fmt.Println("***人間の飲む処理***")
	DrinkAll(eat)

	eat = cheloniaMydas // Turtle型がインターフェースであるEater型に変換される（Turtle型にはインターフェイスが実装されている）
	fmt.Println("***カメ***")
	fmt.Println("***カメの食べる処理***")
	EatAll(eat)
	fmt.Println("***カメの飲む処理***")
	DrinkAll(eat)

}