package arztpraxis

type Medikament struct {
	Name                 string
	Unvertraeglichkeiten []Medikament
}

// AddUnvertraeglichkeit fügt ein anderes Medikament zur Liste der Unverträglichkeiten von m hinzu.
func (m *Medikament) AddUnvertraeglichkeit(other Medikament) {
	m.Unvertraeglichkeiten = append(m.Unvertraeglichkeiten, other)
}

// Prüft, ob m mit other verträglich ist.
func (m Medikament) IstVertraeglich(other Medikament) bool {
	for _, u := range m.Unvertraeglichkeiten {
		if u.Name == other.Name {
			return false
		}
	}
	return true
}
