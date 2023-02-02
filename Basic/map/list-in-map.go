// Mapのメンバーに配列を持たせるパターン
// Map[i][j]みたいな値参照が可能

package main

import "fmt"

type valueList []string

func main(){

	keyList := [5]string{"key1", "key2", "key3","key4","key5"}
	values := [10]string{"value1", "value2", "value3","value4","value5", "value6", "value7","value8","value9","value10"}

	// Map := make(map[string]valueList, len(keyList))
	Map := make(map[string]valueList)
	
	for _, key := range keyList {
		Map[key] = make(valueList, len(values))
		for i,v := range values {
			Map[key][i] = v
		}
	}

	// 結果出力
	for i := range Map {
		fmt.Println("Map[", i, "] = ", Map[i])
	}
	fmt.Println("Map[key2][3] = ", Map["key2"][3])
	
}