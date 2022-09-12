package main

import "fmt"

func main() {
	m := make(map[string]int) // Key=string,Value=intのmapを定義

	m["Foo"] = 10
	fmt.Println("m[\"Foo\"]=", m["Foo"])

	// 更新
	m["Foo"] = 20
	fmt.Println("m[\"Foo\"]=", m["Foo"])

	elem, ok := m["Foo"] // Valueの表示とKeyの存在確認
	fmt.Println("elem:", elem, "ok:", ok)
	elem, ok = m["Bar"]
	fmt.Println("elem:", elem, "ok:", ok)

	// 削除
	delete(m, "Foo")
	fmt.Println("m[\"Foo\"]=", m["Foo"])

	elem, ok = m["Foo"]
	fmt.Println("elem:", elem, "ok:", ok)

}
