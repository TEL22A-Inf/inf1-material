package intro

import "fmt"

type Kunde struct {
	Name    string
	Vorname string
	Kdnr    string
	Punkte  int
}

func Demo1() {
	var k1 Kunde

	// Das leere Struct ausgeben:
	fmt.Println(k1)

	// Felder (Attribute) besetzen
	k1.Name = "Mustermann"
	k1.Vorname = "Max"

	// Das Struct wieder ausgeben:
	fmt.Println(k1)

	k1.Kdnr = "12345-A"
	k1.Punkte = 38

	// Das Struct wieder ausgeben:
	fmt.Println(k1)

	// Vorname ausgeben:
	fmt.Println(k1.Vorname)

	k2 := Kunde{
		Name:    "Beispiel",
		Punkte:  42,
		Vorname: "Barbara",
		Kdnr:    "465b",
	}

	fmt.Println(k2)
}
