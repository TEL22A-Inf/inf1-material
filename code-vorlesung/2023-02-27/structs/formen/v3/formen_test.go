package formen

import "fmt"

func ExampleQuadrat_Draw() {
	q1 := Quadrat{A: 4}
	r1 := Rechteck{A: 4, B: 15}

	q1.Draw()
	fmt.Println()
	r1.Draw()

	q1.Resize(2)
	r1.Resize(2)

	q1.Draw()
	fmt.Println()
	fmt.Println(q1)
	fmt.Println()
	r1.Draw()

	// Output:
}
