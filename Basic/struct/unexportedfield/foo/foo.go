/*
先頭が小文字のフィールドやメソッドは非公開(unexported)と呼ばれ、
package外からアクセスできない
*/

package foo

type Foo struct {
	A int
	B int
	c int // 先頭が小文字なので非公開フィールド(package外からアクセス不可)
	d int // 先頭が小文字なので非公開フィールド(package外からアクセス不可)
}

func (f *Foo) Get_A() int {
	return f.A
}

func (f *Foo) Set_A(v int) {
	f.A = v
}

func (f *Foo) Get_B() int {
	return f.B
}

func (f *Foo) Set_B(v int) {
	f.B = v
}

func (f *Foo) get_c() int { // 先頭が小文字なので非公開メソッド(package外からアクセス不可)
	return f.c
}

func (f *Foo) Set_c(v int) {
	f.c = v
}

func (f *Foo) get_d() int { // 先頭が小文字なので非公開メソッド(package外からアクセス不可)
	return f.d
}

func (f *Foo) Set_d(v int) {
	f.d = v
}