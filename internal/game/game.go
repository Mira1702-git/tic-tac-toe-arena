package game

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Mira1702-git/tic-tac-toe-arena/internal/ai"
	"github.com/Mira1702-git/tic-tac-toe-arena/internal/board"
	"github.com/Mira1702-git/tic-tac-toe-arena/internal/cli"
)

const (
	red   = "\033[31m"
	reset = "\033[0m"
)

type Stats struct {
	XWins int
	OWins int
	Draws int
}

// renderBoard выбирает обычную или большую доску.
func renderBoard(
	b *board.Board,
	config cli.Config,
	winningLine []int,
) string {
	if config.Big {
		return b.RenderBig(
			config.Color,
			winningLine,
		)
	}

	return b.Render(
		config.Color,
		winningLine,
	)
}

// printError печатает ошибку.
// При --color ошибка будет красной.
func printError(config cli.Config, message string) {
	if config.Color {
		fmt.Println(red + message + reset)
		return
	}

	fmt.Println(message)
}

// printStats выводит статистику.
func printStats(
	config cli.Config,
	stats Stats,
	moves int,
) {
	games := stats.XWins +
		stats.OWins +
		stats.Draws

	fmt.Println("=== Stats ===")

	fmt.Printf(
		"Games: %d   %s: %d   %s: %d   Draws: %d\n",
		games,
		config.NameX,
		stats.XWins,
		config.NameO,
		stats.OWins,
		stats.Draws,
	)

	if config.Verbose {
		xRate := 0
		oRate := 0

		if games > 0 {
			xRate = stats.XWins * 100 / games
			oRate = stats.OWins * 100 / games
		}

		fmt.Printf(
			"Moves this game: %d   Win rate — %s: %d%%  %s: %d%%\n",
			moves,
			config.NameX,
			xRate,
			config.NameO,
			oRate,
		)
	}
}

func Run(config cli.Config) {
	stats := Stats{}

	for {
		b := board.New(config.Size)
		current := config.First
		moves := 0

		for {
			// Показываем доску.
			fmt.Print(
				renderBoard(
					b,
					config,
					nil,
				),
			)

			// Ход компьютера в режиме --ai.
			if config.Mode == "ai" && current == "O" {
				index, reason := ai.FindMove(b)

				if index == -1 {
					stats.Draws++
					fmt.Println("Draw!")
					break
				}

				b.Place(index, "O")
				moves++

				if config.Verbose {
					fmt.Printf(
						"AI: %s at %d\n",
						reason,
						index+1,
					)
				}

				fmt.Printf(
					"Computer chooses cell %d\n",
					index+1,
				)

				// Проверяем победу компьютера.
				if b.CheckWin("O") {
					stats.OWins++

					winningLine := b.WinningLine("O")

					fmt.Print(
						renderBoard(
							b,
							config,
							winningLine,
						),
					)

					fmt.Printf(
						"%s wins!\n",
						config.NameO,
					)

					break
				}

				// Проверяем ничью.
				if b.Full() {
					stats.Draws++

					fmt.Print(
						renderBoard(
							b,
							config,
							nil,
						),
					)

					fmt.Println("Draw!")
					break
				}

				current = "X"
				continue
			}

			// Определяем имя текущего игрока.
			playerName := config.NameX

			if current == "O" {
				playerName = config.NameO
			}

			fmt.Printf(
				"%s (%s), choose a cell: ",
				playerName,
				current,
			)

			// Читаем ход через fmt.Scan.
			var cell string
			fmt.Scan(&cell)

			position, err := strconv.Atoi(cell)

			// Ошибка: введено не число
			// или число вне диапазона.
			if err != nil ||
				position < 1 ||
				position > config.Size*config.Size {

				printError(
					config,
					fmt.Sprintf(
						"Error: enter a number 1-%d",
						config.Size*config.Size,
					),
				)

				continue
			}

			index := position - 1

			// Ошибка: клетка уже занята.
			if !b.IsFree(index) {
				printError(
					config,
					fmt.Sprintf(
						"Error: cell %d is taken",
						position,
					),
				)

				continue
			}

			// Ставим X или O.
			b.Place(index, current)
			moves++

			// Проверяем победу.
			if b.CheckWin(current) {
				if current == "X" {
					stats.XWins++
				} else {
					stats.OWins++
				}

				winningLine := b.WinningLine(current)

				fmt.Print(
					renderBoard(
						b,
						config,
						winningLine,
					),
				)

				fmt.Printf(
					"%s wins!\n",
					playerName,
				)

				break
			}

			// Проверяем ничью.
			if b.Full() {
				stats.Draws++

				fmt.Print(
					renderBoard(
						b,
						config,
						nil,
					),
				)

				fmt.Println("Draw!")
				break
			}

			// Меняем игрока.
			if current == "X" {
				current = "O"
			} else {
				current = "X"
			}
		}

		// Статистика после партии.
		printStats(
			config,
			stats,
			moves,
		)

		// Спрашиваем, играть ли ещё.
		for {
			fmt.Print("Play again? (y/n): ")

			var answer string
			fmt.Scan(&answer)

			answer = strings.ToLower(
				strings.TrimSpace(answer),
			)

			if answer == "y" {
				fmt.Println()
				break
			}

			if answer == "n" {
				fmt.Println()
				fmt.Println("Final statistics:")

				printStats(
					config,
					stats,
					moves,
				)

				return
			}

			fmt.Println("Please enter y or n.")
		}
	}
}
