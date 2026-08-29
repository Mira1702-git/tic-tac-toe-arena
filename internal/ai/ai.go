package ai

import "github.com/Mira1702-git/tic-tac-toe-arena/internal/board"

func FindMove(b *board.Board) (int, string) {
	// 1. Попробовать выиграть.
	for i := 0; i < len(b.Cells); i++ {
		if !b.IsFree(i) {
			continue
		}

		b.Place(i, "O")

		if b.CheckWin("O") {
			b.Clear(i)
			return i, "win"
		}

		b.Clear(i)
	}

	// 2. Заблокировать X.
	for i := 0; i < len(b.Cells); i++ {
		if !b.IsFree(i) {
			continue
		}

		b.Place(i, "X")

		if b.CheckWin("X") {
			b.Clear(i)
			return i, "block"
		}

		b.Clear(i)
	}

	// 3. Центр.
	center := len(b.Cells) / 2

	if b.IsFree(center) {
		return center, "center"
	}

	// 4. Углы.
	corners := []int{0, 2, 6, 8}

	for _, index := range corners {
		if b.IsFree(index) {
			return index, "corner"
		}
	}

	// 5. Стороны.
	sides := []int{1, 3, 5, 7}

	for _, index := range sides {
		if b.IsFree(index) {
			return index, "side"
		}
	}

	return -1, ""
}
