package arztpraxis

// Patientendaten
type Patient struct {
	Name string
	// eingenommene Medikamente
	Medis  []Medikament
	Unvies []Medikament // bekannte Unverträglichkeiten
}

// ChecMedi prüft, ob der Patient p das Medikament m einnehmen darf.
func (p Patient) CheckMedi(m Medikament) bool {
	for _, anderesMedi := range p.Medis {
		if !m.IstVertraeglich(anderesMedi) {
			return false
		}
	}
	return true
}
