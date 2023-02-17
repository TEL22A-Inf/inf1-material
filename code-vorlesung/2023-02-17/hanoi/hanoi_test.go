package hanoi

func ExampleBewegePlatte() {
	BewegePlatte("A", "C")

	// Output:
	// A ==> C
}

func ExampleHanoi1() {
	Hanoi1("A", "B", "C")

	// Output:
	// A ==> C
}

func ExampleHanoi() {
	Hanoi(5, "A", "B", "C")

	// Output:
	// A ==> C
}
