package formen

import "fmt"

func ExampleQuadrat_Draw() {
	q1 := Quadrat{A: 4}
	r1 := Rechteck{A: 4, B: 15}

	q1.Draw()
	fmt.Println()
	r1.Draw()

	// Output:
}
