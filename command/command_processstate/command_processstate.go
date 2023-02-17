/*
os.ProcessStateの非公開フィールドにアクセスして更新
*/

package main

import (
	"fmt"
	"os"
	"reflect"
	"unsafe"
)

func main() {
	// var pState *os.ProcessState
	pState := &os.ProcessState{}

	var ref_pState reflect.Value = reflect.ValueOf(pState).Elem()
	var ref_status reflect.Value = ref_pState.FieldByName("status")
	status := (*uint32)(unsafe.Pointer(ref_status.UnsafeAddr()))
	*status = 0

	fmt.Printf("pState: %#v\n", *pState)

}
