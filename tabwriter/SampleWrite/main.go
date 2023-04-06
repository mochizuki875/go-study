package main

import (
	"fmt"
	"os"

	"github.com/liggitt/tabwriter"
)

// kubectlで使ってるやつそのまま
const (
	tabwriterMinWidth = 6
	tabwriterWidth    = 4
	tabwriterPadding  = 3
	tabwriterPadChar  = ' '
	tabwriterFlags    = tabwriter.RememberWidths
)

func main() {

	// witerの作成(kubectlと同じwriterが作られる)
	w := tabwriter.NewWriter(os.Stdout, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, tabwriterFlags)

	// バッファ書き込み(ここをkubectlがリソース情報を書き込むのと同じロジックにしたい)
	// printer.PrintObj(info.Object, w)相当のことをやる

	fmt.Fprintln(w, "a\tb\tc")
	fmt.Fprintln(w, "aa\tbb\tcc")
	fmt.Fprintln(w, "aaa\t")
	fmt.Fprintln(w, "aaaa\tdddd\teeee")

	w.Flush() // ここで書き込んだ内容を出力

}
