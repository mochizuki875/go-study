/*
mmap()による仮想メモリ確保と物理メモリ確保（デマンドページング）の実験
https://pkg.go.dev/syscall#Mmap
https://linuxjm.osdn.jp/html/LDP_man-pages/man2/mmap.2.html
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"time"
)


const(
	ALLOC_SIZE int = 100 * 1024 * 1024 // 100MiB
)


func main(){
	scanner := bufio.NewScanner(os.Stdin)

	PAGE_SIZE := syscall.Getpagesize() // ページサイズを取得(x86_64アーキテクチャなら4096byte)

	fmt.Println("mmap()により仮想メモリを100MiB取得します。(1ページ=" , PAGE_SIZE, "byte)")
	fmt.Println("*Press Enter")
	scanner.Scan() // 入力待ち

	memregion, err := syscall.Mmap(-1, 0, ALLOC_SIZE, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		fmt.Errorf("Faild mmap(): %v", err)
	}

	fmt.Println("mmap()により仮想メモリを100MiB取得しました。")
	fmt.Println("=========================================================")
	fmt.Println("各ページへの1回目の書き込みを行い物理メモリを取得します。(デマンドページング)")
	fmt.Println("*Press Enter")
	scanner.Scan() // 入力待ち

	now := time.Now()

	for i := 0; i<ALLOC_SIZE; i+= PAGE_SIZE {
		memregion[i] = 0
	}

	fmt.Printf("所要時間: %vμs\n", time.Since(now).Microseconds())

	fmt.Println("各ページへの1回目の書き込みおよび物理メモリの取得が完了しました。")
	fmt.Println("*Press Enter")

	scanner.Scan() // 入力待ち

	fmt.Println("=========================================================")
	fmt.Println("各ページへの2回目の書き込みを行います。(1回目の書き込みにより物理メモリは確保済み)")
	fmt.Println("*Press Enter")
	scanner.Scan() // 入力待ち

	now = time.Now()

	for i := 0; i<ALLOC_SIZE; i+= PAGE_SIZE {
		memregion[i] = 0
	}

	fmt.Printf("所要時間: %vμs\n", time.Since(now).Microseconds())

	fmt.Println("各ページへの2回目の書き込みが完了しました。")
	fmt.Println("*Press Enter")
	scanner.Scan() // 入力待ち

}