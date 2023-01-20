package zahlen

// Erwartet eine Zahl n und liefert eine Liste mit den Ziffern von n.
// Annahme: n hat höchstens drei Stellen.
func IntTodigits(n int) []int {
	result := []int{}
	return result
}

// Alternative:

// Erwartet eine Zahl n und liefert sie als String.
func IntToString(n int) string {
	return ""
}

func DigitsToString(digits []int) string {
	digitStrings := []string{
		"null",
		"eins",
		"zwei",
		"drei",
		"vier",
		"fünf",
		"sechs",
		"sieben",
		"acht",
		"neun",
	}

	result := ""
	result += digitStrings[digits[0]]
	result += "hundert"
	result += digitStrings[digits[2]]
	result += "und"
	result += digitStrings[digits[1]]
	result += "zig"

	return result
}
