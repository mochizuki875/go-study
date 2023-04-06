/*
addというサブコマンドを定義
フラグの値を直接変数に格納するパターン

add -t 1,2,3 → 6
add -t 1,2,3 -m → 12
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var target []int
var multi bool

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: add,
}

func add(cmd *cobra.Command, args []string) error {

	sum := 0
	for _, v := range target {
		sum += v
	}

	if multi {
		sum = sum * 2
	}

	fmt.Println(sum)

	return nil
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntSliceVarP(&target, "target", "t", []int{0}, "Help message for target")
	addCmd.Flags().BoolVarP(&multi, "multi", "m", false, "Help message for multi") // -mというフラグがついていたらtrue
}
