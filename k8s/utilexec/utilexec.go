/*
k8sのライブラリで通常のosexec.Commandをwrapして実行
実態としてはosexec.Commandと同じく所定のosコマンドが実行される。
*/

package main

import (
	"fmt"

	utilexec "k8s.io/utils/exec"
)

func main() {

	// k8s.io/utils@v0.0.0-20230209194617-a36077c30491/exec/exec.go

	// &executorをInterfaceに実装したものが返される
	exec := utilexec.New()

	// *cmdWrapper型でosexec.CommandをmaskErrDotCmd()に通したものが返される
	// コマンドのパスが見つからない場合のエラーを&errors.errorStringとして返すため
	command := exec.Command("lls", "-la")
	output, err := command.CombinedOutput()

	fmt.Printf("output: %#v\n", string(output))
	fmt.Printf("err: %#v\n", err)

	fmt.Printf("command: %#v\n", command)
	// fmt.Printf("command: %#v\n", command)
}
