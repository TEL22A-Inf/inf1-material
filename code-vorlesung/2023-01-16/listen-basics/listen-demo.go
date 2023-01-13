package main

import "fmt"

func main() {
	// Eine Variable definieren,
	// in der Zahlen gespeichert werden können.
	var zahl1 int

	// Diese Zahl ausgeben:
	fmt.Println(zahl1)

	// Eine Variable direkt mittels Wert definieren:
	zahl2 := 42
	fmt.Println(zahl2)

	// Eine Variable für eine Liste von Zahlen definieren:
	var list1 []int
	fmt.Println(list1)

	// list1 ist noch leer. Wir hängen Elemente an:
	list1 = append(list1, 55)
	fmt.Println(list1)

	list1 = append(list1, 42, 23, 38)
	fmt.Println(list1)

	// Eine zweite Liste erzeugen:
	list2 := []int{11, 22, 33, 44, 55}
	fmt.Println(list2)

	list1 = append(list1, list2...)
	fmt.Println(list1)

	// Das Element an Stelle 2 von list1 ausgeben:
	fmt.Println(list1[2])

	// Alle Elemente in list1 in einer Schleife ausgeben:
	for i := 0; i < len(list1); i = i + 1 {
		fmt.Println(i, ":", list1[i])
	}

	// Sei i nacheinander jede Position in der Liste list1.
	// Und sei Element jeweils der Wert an dieser Stelle.
	for i, element := range list1 {
		fmt.Println(i, element)
	}

	for i := 0; i < len(list1); i++ {
		element := list1[i]
		fmt.Println(i, element)
	}

	// Jedes Element in list1 verdoppeln:
	for i := 0; i < len(list1); i++ {
		//list1[i] = list1[i] * 2
		list1[i] *= 2
	}
	fmt.Println(list1)

	// Mit range-Schleife.
	// Vorsicht: element darf nur rechts stehen.
	for i, element := range list1 {
		list1[i] = element * 2
		// element *= 2 würde nicht gehen
	}
	fmt.Println(list1)
}
