package main

// Erzeugt eine neue leere Liste von Strings.
func CreateEmptyStringSlice() []string {
	list := []string{}
	return list
}

// Erzeugt eine Liste von Strings mit einem Anfangswert.
func CreateStringSliceWithOneElement(start string) []string {
	return []string{start}
}

// Durchsucht die Liste nach dem Wort w und gibt
// die Position des ersten Vorkommens zurück.
// Wenn w nicht enthalten ist, wird -1 geliefert.
func FindString(list []string, w string) int {

	for i, element := range list {
		if element == w {
			return i
		}
	}

	return -1
}

// Erwartet zwei Listen list1, list2 und ein Wort w.
// Sucht die Position von w in list1 und liefert das
// entsprechende Wort aus list2.
func LookupString(list1, list2 []string, w string) string {
	pos := FindString(list1, w)
	if pos == -1 || pos >= len(list2) {
		return ""
	}
	return list2[pos]

	//return list2[FindString(list1, w)]
}
