package main

import "fmt"

func Foo(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	Foo(n - 1)
	fmt.Println(n)
}

func main() {
	Foo(3)
}
