package main

import "fmt"

func main() {
	// makeで作成
	fmt.Println("makeでmapを作成")
	m1 := make(map[string]int)
	m1["key1"] = 1
	m1["key2"] = 2
	fmt.Println(m1)
	

	// リテラルで生成
	fmt.Println("リテラルでmapを作成")
	m2 := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
		"key4": 4,
		"key5": 5,
	}
	fmt.Println(m2)

	// 値の取得
	fmt.Println("Keyを指定してValueの取得")
	fmt.Println("m2[\"key1\"] = ", m2["key1"])

	fmt.Println("rangeでの取得")
	for key,value := range m2 {
		fmt.Println("m2[\"", key , "\"] = ", value)
	}

	// エントリの追加
	fmt.Println("エントリの追加")
	m2["add_key"] = 100
	fmt.Println(m2)

	// エントリの削除
	fmt.Println("エントリの削除")
	delete(m2, "add_key")
	fmt.Println(m2)

}