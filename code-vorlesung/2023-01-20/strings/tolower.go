package strings

// Erwartet einen String s und liefert
// einen String, bei dem alle Zeichen
// aus s klein geschrieben sind.
func ToLower(s string) string {
	result := ""
	for _, character := range s {
		if character >= 'A' && character <= 'Z' {
			character += ('a' - 'A')
		}
		if character >= 'À' && character <= 'Þ' {
			character += ('a' - 'A')
		}

		result += string(character)
	}
	return result

	// Alternative:
	// return strings.ToLower(s)

}
