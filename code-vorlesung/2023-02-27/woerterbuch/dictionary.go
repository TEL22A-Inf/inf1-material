package woerterbuch

type Dictionary struct {
	entries []Eintrag
}

func (d Dictionary) LookupEn(de string) string {
	for _, element := range d.entries {
		if element.De == de {
			return element.En
		}
	}
	return ""
}

func (d *Dictionary) Add(de, en string) {
	newEntry := Eintrag{de, en}
	// d = append(d, newEntry)
	d.entries = append(d.entries, newEntry)

}
