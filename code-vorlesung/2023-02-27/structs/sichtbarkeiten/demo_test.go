package sichtbarkeiten

import (
	"fmt"

	"github.com/tel22a-inf/inf1-material/code-vorlesung/2023-02-27/structs/intro"
)

func ExampleDemo_private() {
	var k1 intro.Kunde
	k1.Name = "Foo"
	k1.Kdnr = "12345a"

	fmt.Println(k1)

	// Output:

}
