package spielfeld

// Liefert true, falls die angegebene Zeile in board mit
// symbol gefüllt ist.
func RowFull(board [][]string, rowNumber int, symbol string) bool {
	for _, character := range board[rowNumber] {
		if character != symbol {
			return false
		}
	}
	return true
}

// Liefert true, falls irgendeine Zeile in board mit
// symbol gefüllt ist.
func AnyRowFull(board [][]string, symbol string) bool {
	for i := range board {
		if RowFull(board, i, symbol) {
			return true
		}
	}
	return false
}

func PlayerWins(board [][]string, symbol string) bool {
	// TODO
	return true
}
