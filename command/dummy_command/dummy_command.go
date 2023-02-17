/*
	ダミーの*exec.Cmdを作成
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"unsafe"
)

func main() {
	// os.ProcessStateのフィールドは非公開フィールドなのでreflectを使ってアクセス
	pState := &os.ProcessState{}
	var ref_pState reflect.Value = reflect.ValueOf(pState).Elem()
	var ref_status reflect.Value = ref_pState.FieldByName("status")
	status := (*uint32)(unsafe.Pointer(ref_status.UnsafeAddr()))
	*status = 0

	command := &exec.Cmd{ProcessState: pState}

	fmt.Printf("command: %#v\n", command)
	fmt.Println("command.ProcessState.Success()", command.ProcessState.Success())

}
