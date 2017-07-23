package main

import "fmt"

// Vertex 構造体
type Vertex struct {
	X, Y int
}

var (
	// 両方とも宣言
	v1 = Vertex{1, 2}
	// Xのみ宣言
	v2 = Vertex{X: 1}
	// 両方共未宣言
	v3 = Vertex{}
	// ポインタ設定
	p = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
