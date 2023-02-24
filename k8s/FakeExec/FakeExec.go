/*
testexecのFakeCmdを使うことでコマンド実行結果を任意の出力とエラーにエミュレート
*/

package main

import (
	"errors"
	"fmt"

	utilexec "k8s.io/utils/exec"
	testexec "k8s.io/utils/exec/testing"
)

func makeFakeCommandAction(stdout string, err error, cmdFn func()) testexec.FakeCommandAction {
	c := testexec.FakeCmd{
		CombinedOutputScript: []testexec.FakeAction{
			func() ([]byte, []byte, error) {
				if cmdFn != nil {
					cmdFn()
				}
				return []byte(stdout), nil, err
			},
		},
	}
	return func(cmd string, args ...string) utilexec.Cmd {
		return testexec.InitFakeCmd(&c, cmd, args...)
	}
}

func main() {

	fakeCommandAction := makeFakeCommandAction("any output of command.", errors.New("any error"), nil)

	fakeexec := &testexec.FakeExec{
		LookPathFunc: func(s string) (string, error) {
			return "fake-umount", nil
		},
		CommandScript: []testexec.FakeCommandAction{fakeCommandAction},
	}

	// ここからが実際のfakeコマンド実行
	execCmd(fakeexec)
}

// &testexec.FakeExecをutilexec.Interfaceとして受け取りコマンドを実行
func execCmd(exec utilexec.Interface) {

	// &FakeCmdを作成してexec.Cmdに実装
	// command := exec.Command("uname", "-r")
	command := exec.Command("umount", "/tmp/test")

	output, err := command.CombinedOutput()

	fmt.Printf("output: %#v\n", string(output))
	fmt.Printf("err: %#v\n", err)

	// fmt.Printf("command: %#v\n", command)
}
