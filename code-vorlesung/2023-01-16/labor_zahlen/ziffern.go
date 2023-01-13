package laborzahlen

// Erwartet eine Zahl n und liefert eine
// Liste mit den Ziffern dieser Zahl.
/*
func Digits(n int) []int {
	result := []int{}

	for n != 0 {
		lastDigit := n % 10
		result = append([]int{lastDigit}, result...)
		n = n / 10
	}

	return result
}
*/

func Digits(n int) []int {
	if n == 0 {
		return []int{}
	}
	return append(Digits(n/10), n%10)
}

// Liefert die String-Repräsentation von n.
func ToString(n int) string {
	result := ""

	for _, element := range Digits(n) {
		//result += fmt.Sprintf("%d", element)
		result += string('0' + byte(element))
	}
	return result
}
