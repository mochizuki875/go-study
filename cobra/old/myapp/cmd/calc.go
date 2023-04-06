/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// calcCmd represents the calc command
var (
	// Flags
	a int
	b int

	calcCmd = &cobra.Command{
		Use:   "calc",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("calc called")

			// フラグの取得
			a, _ = cmd.Flags().GetInt("a")
			b, _ = cmd.Flags().GetInt("b")

			fmt.Println("a = ", a)
			fmt.Println("b = ", b)

			fmt.Println("a + b = ", a+b)
		},
	}
)

func init() {
	rootCmd.AddCommand(calcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// calcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// calcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// フラグの定義
	calcCmd.Flags().Int("a", 0, "First Value a")
	calcCmd.Flags().Int("b", 0, "Second Value b")

	// viper.BindPFlag("a", calcCmd.Flags().Lookup("a"))
	// viper.BindPFlag("b", calcCmd.Flags().Lookup("b"))
}
