package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	// s[:2]はs[0:2]と同じ
	s = s[:2]
	fmt.Println(s)

	// s[1:]はs[1:0]と同じ
	s = s[1:]
	fmt.Println(s)
}
