package main

import "fmt"

func main() {
	i, j := 42, 2701

	// ポインタ設定
	p := &i
	// ポインタからiを呼び出す。
	fmt.Println(*p)
	// ポインタを通してiに値を設定する。
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
