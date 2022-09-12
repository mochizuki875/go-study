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
}

type JobType struct {
	JobName   string `json:"jobname"`
	JobNumber int    `json:"nobnum"`
}

func main() {
	name := "Mori"
	number := 3
	// jobType := JobType{"worker", 1}
	emp := Employee{
		Name:   name,
		Number: &number,
		// JobType: jobType,
	}

	fmt.Println(emp)
	fmt.Println(reflect.TypeOf(emp))

	emp_json, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(emp_json))
	fmt.Println(reflect.TypeOf(emp_json))

}
