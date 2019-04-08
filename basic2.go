package main

import(
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// for同様丸括弧はいらない。中括弧のみ
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// ifステートメントはforみたいに、条件の前に評価するための簡単なステートメントを書くとこができる
	// 宣言された変数はifスコープ内でのみ有効。else内でも有効
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

// Loopの練習
func Sqrt(x float64) float64 {
	z := float64(1)
	count := 0
	for ; count < 10; count++ {
		z -= (z*z - x) / (2*z)
	}
	fmt.Println(count, " loop")
	return z
}

func main() {
	// Goのforループには丸括弧がいらない。中括弧は必要
	// 初期化ステートメント、条件式、処理後ステートメントをセミコロンで分ける
	sum1 := 0
	for i := 0; i < 10; i++ {
		sum1 += i
	}
	fmt.Println(sum1)

	// 初期化と後処理ステートメントの記述は任意
	sum2 := 1
	for ; sum2 < 1000; {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// セミコロンも省略できちゃう。つまりGoのwhile はfor
	sum3 := 1
	for sum3 < 1000 {
		sum3 *= 3
	}
	fmt.Println(sum3)

	// なんなら条件式まで省略できるから、無限ループにできる（コメントアウトしておくけど）
	// for {
	// }

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println(Sqrt(25))
}
