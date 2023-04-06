/*
greetというサブコマンドの定義
フラグを取るパターン
フラグの値によって処理を分岐する

greet -w morining → Good Morining
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: greet,
}

func greet(cmd *cobra.Command, args []string) {

	when, _ := cmd.Flags().GetString("when") // フラグから値取得

	switch when {
	case "morining":
		fmt.Println("Good Morining")
	case "afternoon":
		fmt.Println("Good Afternoon")
	case "night":
		fmt.Println("Good Night")
	default:
		fmt.Println("Hello World")
	}
}

func init() {
	rootCmd.AddCommand(greetCmd)

	// greetサブコマンドのフラグ定義
	// 第1引数: フラグ名、
	// 第2引数: 省略したフラグ名
	// 第3引数: デフォルト値
	// 第4引数: フラグの説明(Helpで表示)
	greetCmd.Flags().StringP("when", "w", "morining", "Please chose morining/afternoon/night")
}
