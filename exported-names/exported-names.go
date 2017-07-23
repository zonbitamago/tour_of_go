package main

import (
	"fmt"
	"math"
)

func main() {
	// piは外部参照不可のため、小文字になっている。
	// fmt.println(math.pi)

	// 外部参照可能の場合は、必ず大文字になる。
	fmt.Println(math.Pi)
}
