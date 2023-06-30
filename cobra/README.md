# Cobra

spf13/cobra
https://github.com/spf13/cobra#positional-and-custom-arguments

## インストール
`cobra-cli`と`cobra`のインストール
```
$ go install github.com/spf13/cobra-cli@latest
$ go get -u github.com/spf13/cobra@latest
```

## 初期化
コマンドの雛形作成
`--viper=false`は「設定ファイルを使用しない」
```
$ ~/go/bin/cobra-cli init --license MIT --viper=false  
```

`rootCmd = &cobra.Command`ってやつがコマンドの大元。

`cmd/root.go`
```go
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
```

`main.go`では以下のように定義されているので、特にOPやサブコマンドを指定せずにコマンドを実行(=`main.go`をそのまま実行)した場合は`rootCmd`に対して`Execute()`メソッドが実行される。

`main.go`
```go
func main() {
	cmd.Execute()
}
```

## サブコマンドの追加

`version`というサブコマンドを追加。
```
$ ~/go/bin/cobra-cli add version
```

`cmd/version.go`というファイルが追加される。
やっていることは大きく2つ。


①rootCmdへのサブコマンド追加。
```go
func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
```

②サブコマンドにおける処理の定義。
`Run`の部分で記載した処理がサブコマンド実行時に実行される。
```go
// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version called")
	},
}
```

## コマンド引数
サブコマンドに渡したコマンド引数はargsに格納されるので、それをそのまま使えば良い。

`cmd/hello.go`
```go
// サブコマンドに渡したコマンド引数はargsに格納される
func hello(cmd *cobra.Command, args []string) error {
	fmt.Println("Hello " + args[0])

	return nil
}
```

## 参考

【Golang】1日でGoのcobraでサクッとCLIが作れちゃった話  
https://zenn.dev/tama8021/articles/22_0627_go_cobra_cli

