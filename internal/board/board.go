package board

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	reset      = "\033[0m"
	brightRed  = "\033[91m"
	brightBlue = "\033[94m"
	dim        = "\033[2m"
	greenBold  = "\033[1;32m"
)

type Board struct {
	Size  int
	Cells []string
}

func New(size int) *Board {
	return &Board{
		Size:  size,
		Cells: make([]string, size*size),
	}
}

func (b *Board) Reset() {
	for i := range b.Cells {
		b.Cells[i] = ""
	}
}

func (b *Board) IsFree(index int) bool {
	return index >= 0 &&
		index < len(b.Cells) &&
		b.Cells[index] == ""
}

func (b *Board) Place(index int, mark string) bool {
	if !b.IsFree(index) {
		return false
	}

	b.Cells[index] = mark
	return true
}

func (b *Board) Clear(index int) {
	if index >= 0 && index < len(b.Cells) {
		b.Cells[index] = ""
	}
}

func (b *Board) Full() bool {
	for _, cell := range b.Cells {
		if cell == "" {
			return false
		}
	}

	return true
}

func (b *Board) CheckWin(mark string) bool {
	return b.WinningLine(mark) != nil
}

func (b *Board) WinningLine(mark string) []int {
	// Проверяем строки.
	for row := 0; row < b.Size; row++ {
		line := []int{}
		win := true

		for col := 0; col < b.Size; col++ {
			index := row*b.Size + col

			if b.Cells[index] != mark {
				win = false
				break
			}

			line = append(line, index)
		}

		if win {
			return line
		}
	}

	// Проверяем столбцы.
	for col := 0; col < b.Size; col++ {
		line := []int{}
		win := true

		for row := 0; row < b.Size; row++ {
			index := row*b.Size + col

			if b.Cells[index] != mark {
				win = false
				break
			}

			line = append(line, index)
		}

		if win {
			return line
		}
	}

	// Проверяем главную диагональ.
	line := []int{}
	win := true

	for i := 0; i < b.Size; i++ {
		index := i*b.Size + i

		if b.Cells[index] != mark {
			win = false
			break
		}

		line = append(line, index)
	}

	if win {
		return line
	}

	// Проверяем побочную диагональ.
	line = []int{}
	win = true

	for i := 0; i < b.Size; i++ {
		index := i*b.Size + (b.Size - 1 - i)

		if b.Cells[index] != mark {
			win = false
			break
		}

		line = append(line, index)
	}

	if win {
		return line
	}

	return nil
}

// containsIndex проверяет,
// находится ли клетка в победной линии.
func containsIndex(line []int, target int) bool {
	for _, index := range line {
		if index == target {
			return true
		}
	}

	return false
}

// colorValue добавляет цвет к значению клетки.
func colorValue(
	value string,
	display string,
	index int,
	color bool,
	winningLine []int,
) string {
	if !color {
		return display
	}

	// Победная линия имеет приоритет.
	if containsIndex(winningLine, index) {
		return greenBold + display + reset
	}

	if value == "X" {
		return brightRed + display + reset
	}

	if value == "O" {
		return brightBlue + display + reset
	}

	// Пустые клетки содержат номера.
	return dim + display + reset
}

// Render рисует обычную доску.
func (b *Board) Render(
	color bool,
	winningLine []int,
) string {
	var result strings.Builder

	// Определяем ширину самого большого номера.
	// Например:
	// 3×3  → максимум 9  → ширина 1
	// 4×4  → максимум 16 → ширина 2
	// 10×10 → максимум 100 → ширина 3
	contentWidth := len(
		strconv.Itoa(b.Size * b.Size),
	)

	// Полная ширина клетки:
	// пробел + содержимое + пробел.
	cellWidth := contentWidth + 2

	for row := 0; row < b.Size; row++ {
		for col := 0; col < b.Size; col++ {
			index := row*b.Size + col

			value := b.Cells[index]

			if value == "" {
				value = strconv.Itoa(index + 1)
			}

			// Выравниваем содержимое клетки вправо
			// до одинаковой ширины.
			display := fmt.Sprintf(
				"%*s",
				contentWidth,
				value,
			)

			display = colorValue(
				value,
				display,
				index,
				color,
				winningLine,
			)

			result.WriteString(" ")
			result.WriteString(display)
			result.WriteString(" ")

			if col < b.Size-1 {
				result.WriteString("|")
			}
		}

		result.WriteString("\n")

		if row < b.Size-1 {
			for col := 0; col < b.Size; col++ {
				result.WriteString(
					strings.Repeat("-", cellWidth),
				)

				if col < b.Size-1 {
					result.WriteString("+")
				}
			}

			result.WriteString("\n")
		}
	}

	return result.String()
}

// RenderBig рисует большую доску.
// Каждая клетка имеет размер 5x3.
func (b *Board) RenderBig(
	color bool,
	winningLine []int,
) string {
	var result strings.Builder

	for row := 0; row < b.Size; row++ {
		// Каждая клетка состоит из трёх строк.
		for glyphRow := 0; glyphRow < 3; glyphRow++ {
			for col := 0; col < b.Size; col++ {
				index := row*b.Size + col

				lines := b.bigCell(
					index,
					color,
					winningLine,
				)

				result.WriteString(lines[glyphRow])

				if col < b.Size-1 {
					result.WriteString("|")
				}
			}

			result.WriteString("\n")
		}

		// Горизонтальная граница между рядами.
		if row < b.Size-1 {
			for col := 0; col < b.Size; col++ {
				result.WriteString("-----")

				if col < b.Size-1 {
					result.WriteString("+")
				}
			}

			result.WriteString("\n")
		}
	}

	return result.String()
}

// bigCell возвращает три строки одной большой клетки.
func (b *Board) bigCell(
	index int,
	color bool,
	winningLine []int,
) [3]string {
	value := b.Cells[index]

	// Клетка с X.
	if value == "X" {
		lines := [3]string{
			"X   X",
			"  X  ",
			"X   X",
		}

		return colorBigCell(
			lines,
			"X",
			index,
			color,
			winningLine,
		)
	}

	// Клетка с O.
	if value == "O" {
		lines := [3]string{
			" OOO ",
			"O   O",
			" OOO ",
		}

		return colorBigCell(
			lines,
			"O",
			index,
			color,
			winningLine,
		)
	}

	// Пустая клетка.
	number := strconv.Itoa(index + 1)

	lines := [3]string{
		"     ",
		centerBig(number),
		"     ",
	}

	if color {
		lines[1] = dim + lines[1] + reset
	}

	return lines
}

// colorBigCell окрашивает большой X или O.
func colorBigCell(
	lines [3]string,
	mark string,
	index int,
	color bool,
	winningLine []int,
) [3]string {
	if !color {
		return lines
	}

	code := ""

	if containsIndex(winningLine, index) {
		code = greenBold
	} else if mark == "X" {
		code = brightRed
	} else {
		code = brightBlue
	}

	for i := 0; i < 3; i++ {
		lines[i] = code + lines[i] + reset
	}

	return lines
}

// centerBig помещает номер в центр клетки шириной 5.
func centerBig(value string) string {
	padding := 5 - len(value)

	if padding <= 0 {
		return value
	}

	left := padding / 2
	right := padding - left

	return strings.Repeat(" ", left) +
		value +
		strings.Repeat(" ", right)
}
