package main

import(
	"fmt"
	"strings"
	"math"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

type Vertex struct {
	// 構造体はフィールドの集まり
	X int
	Y int
}

// スライスの練習
// 長さ dy のsliceに、各要素が8bitのunsigned int型で長さ dx のsliceを割り当てたものを返すように実装する
// go tourのターミナルだと画像がみれた
func Pic(dx, dy int) [][]uint8 {
	picY := make([]uint8, dy, dy)
	picX := make([][]uint8, dx, dx)
	for i := range picX {
		for j := range picY {
			picY[j] = uint8((i + j) / 2)
		}
		picX[i] = picY
	}
	return picX
}

// mapの練習
// 文章に各単語か何回出現するかをカウントする
func WordCount(s string) map[string]int {
	wordArr := strings.Fields(s)
	wordMap := make(map[string]int)
	for _, w := range wordArr {
		_, ok := wordMap[w]
		if ok {
			wordMap[w]++
		} else {
			wordMap[w] = 1
		}
	}
	return wordMap
}

// 関数も変数。引数にも戻り値にも使えます
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// クロージャの練習
// フィボナチ数列を返すクロージャを返す
func fibonacci() func() int {
	fibArr := []int{0, 1}
	return func() int {
		ret := fibArr[0]
		fibArr = append(fibArr, fibArr[0] + fibArr[1])[1:]
		return ret
	}
}

func main() {
	// Goにもポインタがあるぞ
	// 型Tのポインタはの型は*T型でゼロ値はnil
	number := 1024
	pn := &number

	// ポインタを利用して変数を読み出す場合*オペレータを使う
	fmt.Println(*pn)
	*pn = 64
	fmt.Println(*pn)

	// 構造体のアクセスは.を用いる
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)

	// フィールドにはポインタを使ってアクセスすることもできる
	// 構造体vのポインタpがあるとすれば、
	// vのフィールドXにアクセスするには(*p).Xと書くのが通常の型と同じ書き方
	// Goではp.Xでアクセスできる
	p := &v
	p.Y = 1
	fmt.Println(v)

	// Goの配列
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "World"
	fmt.Println(arr)
	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes)

	// []Tは型Tのスライス表記
	// ゼロ値（var s []int）はnil
	// a[1:3]はaの１,2の要素を持つ配列
	var partPrimes []int = primes[1:3]
	fmt.Println(partPrimes)

	// スライスは配列の参照みたいなもの
	// スライスを変更すると元配列にも影響あり。その逆もあり
	partPrimes[1] = 10
	fmt.Println(primes)
	primes[1] = 100
	fmt.Println(partPrimes)

	// スライスの上限下限
	// Pythonっぽいぞー
	fmt.Println(primes[3:])
	fmt.Println(primes[:3])
	fmt.Println(primes[:])

	// スライスには長さと容量がある
	// スライスの始まるところによって容量が変わる
	fmt.Printf("len([:0]) = %d, cap([:0]) = %d\n", len(primes[:0]), cap(primes[:0]))
	fmt.Printf("len([2:]) = %d, cap([2:]) = %d\n", len(primes[2:]), cap(primes[2:]))
	fmt.Printf("len([1:4]) = %d, cap([1:4]) = %d\n", len(primes[1:4]), cap(primes[1:4]))

	// スライスは組み込み関数のmakeで作成することができる
	// a := make(スライスの型, 長さ, 容量)
	sli := make([]int, 0, 5)
	fmt.Printf("%d, %d\n", len(sli), cap(sli))

	// スライスに要素を追加するときは組み込み関数のappendを使う
	// append(追加元のスライス, 追加したい変数群)
	// 元スライスの容量を超える場合より大きな配列を割り当てるが、別の割当先を参照している
	sli = append(sli, 0)
	fmt.Println(sli)
	sli = append(sli, 1, 2, 3, 4, 5, 6)
	fmt.Println(sli)

	// rangeは配列をforで回すときに用いる
	// 戻り値の１つ目がインデックス、２つ目が要素のコピー
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// rangeでインデックスが必要ない場合 _ にして捨てることができる
	// インデックスだけ必要な場合はvalueは書く必要がない
	for i := range pow {
		pow[i] = 1 << (uint(i) + 1)
	}
	for _, v := range pow {
		fmt.Println(v)
	}

	pic.Show(Pic)

	// mapはキーと値を関連づける
	// ゼロ値はnilで、キーを持たずまた追加もできないのでmakeで型を指定して初期化されたものを使う
	// make(map[キーの型]値の型)
	// makeを使わない場合
	// var m = map[string]Vertex{"Hiroto": {5, 14}}
	birthdayMap := make(map[string]Vertex)
	birthdayMap["Naoki"] = Vertex{10, 30}
	fmt.Println(birthdayMap["Naoki"])

	//  要素の挿入や更新
	// 要素の取得
	birthdayMap["Hiroto"] = Vertex{5, 14}
	fmt.Println(birthdayMap["Hiroto"])
	// 要素の削除
	delete(birthdayMap, "Naoki")
	// キーに対して要素が存在するかは取得の第２戻り値
	_, ok := birthdayMap["Naoki"]
	fmt.Println(ok)

	wc.Test(WordCount)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
