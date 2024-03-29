/*
大元になるコマンド
単純にコマンドを単体実行するとrootCmd.Execute()が実行される
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// 大元になるコマンド
// 単純にコマンドを単体実行するとrootCmd.Execute()が実行される

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// ここでRunを定義すればコマンドを単体実行した時の処理を定義できる
	// Runを定義していない場合はHelpメッセージとしてLongの文言が返される

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },

	Run: runHelp, // cmd.Help()を呼び出すとrootCmdのLong+サブコマンド+フラグがヘルプとして表示される(なくても同じ)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.myapp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
