package spielfeld

import "fmt"

func ExampleAnyRowFull() {
	b1 := [][]string{
		{" ", "X", " "},
		{"X", "X", "X"},
		{" ", "X", " "},
	}

	fmt.Println(AnyRowFull(b1, "X"))
	fmt.Println(AnyRowFull(b1, "O"))

	// Output:
	// true
	// false
}
