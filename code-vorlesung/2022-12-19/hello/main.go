package main

import "fmt"

func main() {

	// Eine erste Ausgabe.
	fmt.Println("Gib ein Wort ein: ")

	// Ein Wort einlesen.
	var input string
	fmt.Scanln(&input)

	// Dieses Wort wieder ausgeben.
	fmt.Println("Du hast ", input, " eingegeben.")
}
