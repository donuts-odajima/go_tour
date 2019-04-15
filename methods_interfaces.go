package main

import(
	"fmt"
	"math"
	"io"
	"os"
	"strings"
	"golang.org/x/tour/reader"
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Vertex struct {
	X, Y float64
}

type ErrNegativeSqrt float64

type MyReader struct{}

// 自分のImageメソッドを作る
// "golang.org/x/tour/pic.ShowImage"で表示できるように
// imageインターフェースを満たすようメソッドを作る
type Image struct{}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

// Readerの練習
// 'A' の無限ストリームを出力するReader型を実装する
// 評価は"golang.org/x/tour/reader" のValidationメソッドが行う
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

// Readerの練習２
// ROT13という形式で暗号化された文字を展開するReaderを実装する
type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := 0; i < n; i++ {
		switch {
			case (b[i] >= 'A' && b[i] < 'N') || (b[i] >= 'a' && b[i] < 'n'):
				b[i] += 13
			case (b[i] >= 'N' && b[i] <= 'Z') || (b[i] >= 'n' && b[i] <= 'z'):
				b[i] -= 13
		}
	}
	return n, err
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		return math.Sqrt(x), nil
	} else {
		return 0, ErrNegativeSqrt(x)
	}
}

// Stringerの練習　IPアドレスを文字列に
type IPAddr [4]byte
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// Goにはクラスがない。が、型にメソッドを定義できる
// レシーバという特殊な引数を関数にとることで、その関数をレシーバで指定した型のメソッドとする
// 型ExampleにEx1(x, y, int) int {}というメソッドを追加する場合
// func (e Example) Ex1(x, y, int) int {}
// レシーバを伴う関数の宣言はレシーバ型が同じパッケージにある必要がある
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// メソッドの元のレシーバ変数を変更したいときは、ポインタレシーバを使用する
// v := Vertex{3, 4}
// v.Scale(10) は 内部で(&v).Scale(10) となっているらしい
// ポインタレシーバは変数のコピーを避けることに貢献している
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

// インターフェースはルールみたいなもの
// 指定した名前のメソッドを含む型の変数のみ宣言できる
type Areaer interface {
	Area()
}

type Circle struct {
	Radius float64
}

type Trapezoid struct {
	Top, Bottom, Height float64
}

func (c *Circle) Area() {
	// インターフェース自体の中にある具体的な値がnilの場合、レシーバはnilをレシーバとして呼び出す
	// そのためnilが来てもいいよいにメソッドを記述するのが普通らしい
	// インターフェース自体はnilでないので注意
	if c == nil {
		fmt.Println("<nil>")
		return
	}
	c.Radius += 1
	fmt.Println(c.Radius*c.Radius*math.Pi)
}

func (t *Trapezoid) Area() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println((t.Top + t.Bottom) * t.Height / 2)
}

func typePrint(i interface{}) {
	// 型スイッチ
	// そのまんま、型でスイッチできる
	switch v := i.(type) {
	case int:
		fmt.Printf("int %v\n", v)
	case float64:
		fmt.Printf("float64 %v\n", v)
	default:
		fmt.Println("i dont know...")
	}
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

	var i Areaer
	// nilなままのインターフェースは値も具体的な型ももたない
	// 呼び出さないように！ランタイムエラーになる
	i = &Circle{3}
	i.Area()
	i = &Trapezoid{ Top: 2, Bottom: 3, Height: 4 }
	i.Area()

	// 空のインターフェースは未知の値を保持できる
	// なぜなら、空のインターフェースはゼロ個のメソッドを指定したと考えることができ、
	// 全ての型は少なくともゼロ個のメソッドを実装しているため
	var valiableI interface{}
	valiableI = 42
	fmt.Println(valiableI)
	valiableI = "hello"
	fmt.Println(valiableI)

	// 型アサーション
	// ある変数iの値が型Tを満たしてるか（j と定義できるか）
	var i2 interface{} = "hello"
	j, ok := i2.(string)
	fmt.Println(j, ok)
	k, ok := i2.(float64)
	fmt.Println(k, ok)

	typePrint(10)
	typePrint(1.23)
	typePrint("a")

	hosts := map[string]IPAddr{
		"naoki": {1, 2, 3, 4},
		"donuts": {4, 3, 2, 1},
	}
	for name, host := range hosts {
		fmt.Printf("%v: %v\n", name, host)
	}

	fmt.Println(Sqrt(-2))
	fmt.Println(Sqrt(2))

	r := strings.NewReader("Hello, World!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	reader.Validate(MyReader{})


	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r2 := rot13Reader{s}
	io.Copy(os.Stdout, &r2)

	img := Image{}
	pic.ShowImage(img)
}
