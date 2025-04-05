/*
  field.Path
  良い感じにパスっぽい文字列(AAA.BBB.CCC)を生成したり扱うことができるクラス
  インデックスを付与したりとかできるが基本的に文字列以上の意味はない
*/

package main

import (
	"fmt"
	"strings"

	"k8s.io/utils/field"
)

func main() {

	fldPath := field.NewPath("spec")

	// field.PathではString()メソッドが定義されているため、インスタンス名を指定すれば文字列でパスを取得できる
	fmt.Println("fldPath: ", fldPath) // => spec

	new_path_1 := fldPath.Child("topologySpreadConstraints") // fldPath.Child("topologySpreadConstraints")で"fldPath"に"topologySpreadConstraints"を追加した新しいPathを作成
	fmt.Println("new_path_1: ", new_path_1)                  // => spec.topologySpreadConstraints
	fmt.Println("fldPath: ", fldPath)                        // => spec (元のPathは変わらない)

	fmt.Println("new_path_2.Root(): ", new_path_1.Root()) // => spec(ルートパスを取得)

	new_path_2 := new_path_1.Child("maxSkew")             // new_path_1.Child("maxSkew")で"new_path_1"に"maxSkew"を追加した新しいPathを作成
	fmt.Println("new_path_2: ", new_path_2)               // => spec.topologySpreadConstraints.maxSkew
	fmt.Println("new_path_2.Root(): ", new_path_2.Root()) // => spec(ルートパスを取得)

	for i := 0; i < 3; i++ {
		subPath := new_path_1.Index(i) // Index()を用いることで、Pathにインデックスを追加できる
		fmt.Printf("subPath[%d]: %s \n", i, subPath)
	}

	for i := 0; i < 3; i++ { // Parentのindexに対してchildを付与
		subPath := new_path_1.Index(i)                                       // Index()を用いることで、Pathにインデックスを追加できる
		fmt.Printf("subPath[%d]: %s \n", i, subPath.Child("matchLabelKeys")) // => spec.topologySpreadConstraints[i].matchLabelKeys

	}

	// 特定の文字列がPathに含まれているかを確認
	fldPath = nil
	fldPath = field.NewPath("spec")
	fmt.Println("fldPath: ", fldPath) // => spec
	fldPath = fldPath.Child("template")
	fmt.Println("fldPath: ", fldPath) // => spec.template
	fldPath = fldPath.Child("spec")
	fmt.Println("fldPath: ", fldPath) // => spec.template.spec

	if strings.Contains(fldPath.String(), "spec.template") {
		fmt.Println("OK") // => OK
	}
}
