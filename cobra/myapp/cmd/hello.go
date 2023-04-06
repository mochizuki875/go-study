/*
helloというサブコマンドの定義
コマンド引数を取るパターン

hello World → Hello World!
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// サブコマンドにおける処理の定義
// RunEの部分で記載した処理がサブコマンド実行時に実行される
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: hello,                 // エラーを返す場合はRunEを使う
	Args: cobra.MinimumNArgs(1), // コマンド引数の最低個数を指定したい場合はこのように指定
}

// サブコマンドに渡したコマンド引数はargsに格納される
func hello(cmd *cobra.Command, args []string) error {
	fmt.Println("Hello " + args[0] + "!")

	return nil
}

func init() {
	rootCmd.AddCommand(helloCmd) // rootCmdへのサブコマンド追加
}
