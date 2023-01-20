package laborzahlen

import "fmt"

func ExampleDigits() {
	fmt.Println(Digits(123))
	fmt.Println(Digits(15724))

	// Output:
	// [1 2 3]
	// [1 5 7 2 4]
}

func ExampleToString() {
	fmt.Println(ToString(123))
	fmt.Println(ToString(15724))

	// Output:
	// 123
	// 15724
}
