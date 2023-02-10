package main

import "fmt"

func Add(n, m int) int {
	if m == 0 {
		return n
	}
	return 1 + Add(n, m-1)
}

func main() {
	fmt.Println(Add(3, 2))
}
