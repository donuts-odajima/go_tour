// A Tour of Go

// プログラムはパッケージで構成される。
// プログラムはmainパッケージから開始される
package main

// インポートは括弧でくくると一括でできる（factored import statement）
import (
	// フォーマットパッケージ
	"fmt"
	"time"
	"math"
	// パッケージ名はインポートパスの最後の要素と同じ名前
	"math/rand"
)

// 型名は変数名の後ろに書く（戻り値の型も同様）
// 関数の２つ以上の引数の型が同じなら省略して書ける
func add(x, y int) int {
	return x + y
}

// 戻り値は複数可能
func swap(x, y string) (string, string) {
	return y, x
}

// 戻り値に名前をつけることができる（named return value）
// これによってreturn ステートメントに何も書かずに戻すことができる（naked return）
// ただし長い関数では読みやすさの観点から使うべきではない
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 変数宣言。型の引数同様に最後に書くだけで複数の型を指定できる
var cpp, python, java bool

func main() {
	fmt.Println("Hello, Donuts!")

	fmt.Println("My favorite number is ", rand.Intn(10))

	fmt.Println("The time is ", time.Now())

	// 最初が大文字でで始まる名前は、外部のパッケージからエクスポートされた名前
	// （小文字のものにはアクセスできない）
	fmt.Println("π = ", math.Pi)

	fmt.Println("33 + 4 = ", add(33, 4))

	// swap(17)._1 とか直では値をとれないっぽい
	a, b := swap("hello", "donuts")
	fmt.Println("hello donuts --swap--> ", a, b)

	c, d := split(17)
	fmt.Println("17 --split--> ", c, d)

	// var ステートメントはパッケージ内または関数内でのみ使用できる
	// 初期化せずに変数を宣言するとゼロ値が与えられる
	var scala int
	fmt.Println(scala, cpp, python, java)

	// 初期化子が与えられている場合は型が省略できる
	var iphone, ipad, mac = true, "no!", 1
	fmt.Println(iphone, ipad, mac)

	// 関数の中であればvar の代わりに:= で代入文とすることで暗黙的な型宣言となる
	android := false
	fmt.Println(android)

	// 変数v の型をT に変換したい場合、T(v)
	// var による宣言の場合型を明示的に記述しなくてはならない
	var birth int = 1030
	var str_birth string = string(birth)
}
