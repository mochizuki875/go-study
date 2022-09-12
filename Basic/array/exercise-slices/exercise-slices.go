package main

// import "fmt"
import "golang.org/x/tour/pic"

// uint8型の2次元配列（slice)を作成する関数
func CreateMatrixArray(dx, dy int) [][]uint8 {
	array := make([][]uint8, dy, dy) // sliceを2次元で初期化
	for i, _ := range array {
		array[i] = make([]uint8, dx, dx)
	}
	return array
}

// 画像処理したuint8型の2次元配列（slice)を作成する関数
func Pic(dx, dy int) [][]uint8 {

	pic := CreateMatrixArray(dx, dy)

	// この部分で各配列要素を加工
	for i, _ := range pic {
		for j, _ := range pic[i] {
			pic[i][j] = uint8((i + j) / 2)
		}
	}
	
	return pic
}

func main() {

	// 画像を生成する関数
	pic.Show(Pic)

}