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

func main() {

	// 事前にtestexec.FakeCmdを作ってCombinedOutputScript()実行結果を定義しておく
	c := testexec.FakeCmd{
		CombinedOutputScript: []testexec.FakeAction{
			func() ([]byte, []byte, error) { // FakeActionとしてstdoutとerrを返す関数を定義
				return []byte("any output of command."), nil, errors.New("any error")
			},
		},
	}

	// Commandが実行された場合に上で定義したFakeCmdを返すようにする

	//   fakeCommandActionとして上で定義したFakeCmdをexec.Cmdで返す関数を定義
	fakeCommandAction := func(cmd string, args ...string) utilexec.Cmd {
		return testexec.InitFakeCmd(&c, cmd, args...)
	}

	//   CommandScriptに上で作成したfakeCommandActionを設定することでCommand()が実行されたらFakeCmdをexec.Cmdで返すようにする
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
