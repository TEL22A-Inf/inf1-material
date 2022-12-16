package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMicro())

	// Spieler wählt eine Zahl
	var input int = ReadNumber("Geben Sie eine Zahl zwischen 0 und 36 ein: ")

	// Am Rad drehen
	result := SpinWheel()

	// Auswerten, ob der Spieler gewonnen hat
	PrintResult(input, result)
}

// Fragt den Spieler nach einer Zahl und liefert diese.
func ReadNumber(message string) int {
	i := 0
	for ; i < 10; i = i + 1 {
		fmt.Print(message)

		var number int
		fmt.Scanln(&number)

		if number >= 0 && number <= 36 {
			return number
		}
		fmt.Println("Die Zahl war ungültig!")
	}
	fmt.Println("Idiot!")
	return -1
}

// Prüft, ob der Spieler mit seiner gewählt Zahl gewonnen hat.
// Braucht dazu die gewählte Zahl und die Zahl, die gedreht wurde.
func PlayerWins(playerNumber, wheelNumber int) bool {
	if playerNumber == wheelNumber {
		return true
	} else {
		return false
	}
}

// Gibt das Ergebnis auf die Konsole aus.
// Braucht dazu die gewählte Zahl und die Zahl, die gedreht wurde.
func PrintResult(playerNumber, wheelNumber int) {
	if PlayerWins(playerNumber, wheelNumber) {
		fmt.Println("Gewonnen")
	} else {
		fmt.Println("Verloren, es wurde", wheelNumber, "gedreht")
	}
}

// Dreht das Roulette-Rad und liefert ein Ergebnis.
func SpinWheel() int {
	result := rand.Intn(37)
	return result
}
