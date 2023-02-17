package hanoi

import (
	"fmt"
)

func BewegePlatte(start, ziel string) {
	fmt.Printf("%s ==> %s\n", start, ziel)
}

func Hanoi1(start, mitte, ziel string) {
	BewegePlatte(start, ziel)
}

func Hanoi(h int, start, mitte, ziel string) {
	if h == 0 {
		return
	}
	Hanoi(h-1, start, ziel, mitte)
	BewegePlatte(start, ziel)
	Hanoi(h-1, mitte, start, ziel)
}
