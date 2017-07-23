package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		// ==2**i
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d¥n", value)
	}
}
