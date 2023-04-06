/*
  GoでJsonを扱う
  `json: "hoge"`みたいな構造体タグを付与することでフィールドをJsonで扱える

  https://qiita.com/Syuparn/items/9e6fb68afb5418893c23
*/

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Employee struct {
	Name    string `json:"empname"`
	Number  *int   `json:"empid"`
	JobType `json:"jobtype,omitempty"`
	Member  MemberList `json:"member,omitempty"`
}

type JobType struct {
	JobName   string `json:"jobname"`
	JobNumber int    `json:"nobnum"`
}

type MemberList map[string]string

func main() {
	number := 3
	emp := Employee{
		Name:    "Mori",
		Number:  &number,
		JobType: JobType{"worker", 1},
		Member:  MemberList{"A": "memberA", "B": "memberB", "C": "memberC"},
	}

	fmt.Println("struct = ", emp) // 普通にstructとして出力
	fmt.Println(reflect.TypeOf(emp))

	emp_json, err := json.Marshal(emp) // Jsonとして出力
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("json = ", string(emp_json))
	fmt.Println(reflect.TypeOf(emp_json))

}
