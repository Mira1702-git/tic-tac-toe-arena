# Tic-Tac-Toe Arena

Tic-Tac-Toe Arena is a command-line Tic-Tac-Toe game written in Go.

The game supports:

- two-player mode;
- player vs computer mode;
- colored output;
- large board rendering;
- custom player names;
- configurable first player;
- configurable board size in two-player mode;
- session statistics;
- verbose statistics;
- deterministic AI.

## Run

Run the project from the repository root:

```bash
go run . --players
```

or:

```bash
go run . --ai
```

Exactly one game mode must be selected.

## Modes

### Two players

```bash
go run . --players
```

Two human players take turns playing X and O.

### Player vs AI

```bash
go run . --ai
```

The human player is X and the computer is O.

The AI is available only for the standard 3x3 board.

## Options

```text
--players        two human players take turns
--ai             play against the computer
--color          enable colored output
--big            render the board with large glyphs
--verbose        show extended statistics
--first X|O      choose who moves first (default: X)
--name A,B       set custom names for X and O
--size N         use an N×N board (default: 3, minimum: 3)
--help, -h       show help
```

Examples:

```bash
go run . --players --color
```

```bash
go run . --players --big
```

```bash
go run . --players --big --color
```

```bash
go run . --players --first O
```

```bash
go run . --players --name Alice,Bob
```

```bash
go run . --players --size 4
```

```bash
go run . --players --verbose
```

```bash
go run . --ai --verbose
```

## Rules

The board cells are numbered from `1` to `N²`.

For the default 3x3 board:

```text
 1 | 2 | 3
---+---+---
 4 | 5 | 6
---+---+---
 7 | 8 | 9
```

Players take turns choosing a free cell.

Players use `X` and `O`. The `--first` option determines which mark moves first.

A player wins by filling a complete:

- row;
- column;
- diagonal.

If the board becomes full and nobody has won, the game ends in a draw.

After every game, the program asks:

```text
Play again? (y/n):
```

Entering `y` starts a new game while keeping the session score.

Entering `n` prints the final statistics and exits.

## AI

In `--ai` mode:

- the human is X;
- the computer is O;
- the board is always 3x3.

The AI uses the following deterministic strategy:

1. Win if possible.
2. Block the opponent if necessary.
3. Choose the center.
4. Choose a corner.
5. Choose a side.

Corner priority:

```text
1, 3, 7, 9
```

Side priority:

```text
2, 4, 6, 8
```

The AI does not use minimax.

With `--verbose`, the program also explains the AI choice, for example:

```text
AI: block at 5
```

## Color mode

Use:

```bash
go run . --players --color
```

Color mode uses ANSI terminal colors:

- X — bright red;
- O — bright blue;
- empty cell numbers — dim;
- winning line — green and bold;
- error messages — red.

Without `--color`, the program uses plain text.

## Big mode

Use:

```bash
go run . --players --big
```

Each cell is rendered as a 5x3 block.

Example X:

```text
X   X
  X
X   X
```

Example O:

```text
 OOO
O   O
 OOO
```

Big mode can also be combined with color:

```bash
go run . --players --big --color
```

## Custom board size

Two-player mode supports an N×N board:

```bash
go run . --players --size 4
```

Example:

```text
  1 |  2 |  3 |  4
----+----+----+----
  5 |  6 |  7 |  8
----+----+----+----
  9 | 10 | 11 | 12
----+----+----+----
 13 | 14 | 15 | 16
```

`N` must be at least `3`.

`--size` cannot be combined with `--ai`.

## Custom names

Use:

```bash
go run . --players --name Alice,Bob
```

Then Alice plays X and Bob plays O.

Example:

```text
Alice (X), choose a cell:
```

The custom names are also used in win messages and statistics.

For example:

```text
Alice wins!
```

## Statistics

Statistics are kept for the entire session.

Example:

```text
=== Stats ===
Games: 3   X: 1   O: 1   Draws: 1
```

With custom names:

```text
=== Stats ===
Games: 1   Alice: 1   Bob: 0   Draws: 0
```

With `--verbose`, additional information is displayed:

```text
=== Stats ===
Games: 3   X: 1   O: 1   Draws: 1
Moves this game: 7   Win rate — X: 33%  O: 33%
```

## Example game

Run:

```bash
go run . --players
```

Example:

```text
 1 | 2 | 3
---+---+---
 4 | 5 | 6
---+---+---
 7 | 8 | 9
X (X), choose a cell: 1

 X | 2 | 3
---+---+---
 4 | 5 | 6
---+---+---
 7 | 8 | 9
O (O), choose a cell: 4

 X | 2 | 3
---+---+---
 O | 5 | 6
---+---+---
 7 | 8 | 9
X (X), choose a cell: 2

 X | X | 3
---+---+---
 O | 5 | 6
---+---+---
 7 | 8 | 9
O (O), choose a cell: 5

 X | X | 3
---+---+---
 O | O | 6
---+---+---
 7 | 8 | 9
X (X), choose a cell: 3

 X | X | X
---+---+---
 O | O | 6
---+---+---
 7 | 8 | 9
X wins!
=== Stats ===
Games: 1   X: 1   O: 0   Draws: 0
Play again? (y/n):
```

## Input validation

Invalid input does not crash the game.

For example:

```text
X (X), choose a cell: hello
Error: enter a number 1-9
```

If a cell is already occupied:

```text
Error: cell 1 is taken
```

The program then asks for another move.

Invalid command-line options also produce an error and display usage information.

Examples of invalid configurations include:

- no game mode selected;
- both `--players` and `--ai` selected;
- an unknown flag;
- an invalid value for `--first`;
- a board size smaller than 3;
- using `--size` together with `--ai`;
- invalid or empty player names.

## Project structure

```text
tic-tac-toe-arena/
├── go.mod
├── README.md
├── main.go
└── internal/
    ├── ai/
    │   └── ai.go
    ├── board/
    │   └── board.go
    ├── cli/
    │   └── cli.go
    └── game/
        └── game.go
```

The project is split into separate packages:

- `cli` — command-line argument parsing and validation;
- `board` — board state, win detection and rendering;
- `game` — game loop, input and session statistics;
- `ai` — deterministic computer-player strategy.

## Team

- Adil Tolegen — @github-username
- Gaide Musenova — @github-username
- Merey Zhaxybayeva — @github-username