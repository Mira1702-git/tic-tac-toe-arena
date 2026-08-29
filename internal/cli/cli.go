package cli

import (
	"errors"
	"strconv"
	"strings"
)

type Config struct {
	Mode    string
	Color   bool
	Big     bool
	Verbose bool
	First   string
	NameX   string
	NameO   string
	Size    int
	Help    bool
}

func Usage() string {
	return `Usage: go run main.go (--players | --ai) [options]

Modes (exactly one required):
  --players        two human players take turns
  --ai             play against the computer (you are X)

Options:
  --color          enable colored output (default: plain)
  --big            render the board with large glyphs
  --verbose        show extended statistics
  --first X|O      who moves first (default: X)
  --name A,B       custom names: X=A, O=B
  --size N         board is N×N, win = N in a row (default: 3)
  --help, -h       print this help and exit 0

`
}

func Parse(args []string) (Config, error) {
	config := Config{
		First: "X",
		NameX: "X",
		NameO: "O",
		Size:  3,
	}

	playersMode := false
	aiMode := false
	sizeWasSet := false

	for i := 0; i < len(args); i++ {
		arg := args[i]

		switch arg {
		case "--players":
			playersMode = true

		case "--ai":
			aiMode = true

		case "--color":
			config.Color = true

		case "--big":
			config.Big = true

		case "--verbose":
			config.Verbose = true

		case "--help", "-h":
			config.Help = true
			return config, nil

		case "--first":
			if i+1 >= len(args) {
				return config, errors.New("missing value for --first")
			}

			i++
			config.First = args[i]

			if config.First != "X" && config.First != "O" {
				return config, errors.New("--first must be X or O")
			}

		case "--name":
			if i+1 >= len(args) {
				return config, errors.New("missing value for --name")
			}

			i++

			parts := strings.Split(args[i], ",")

			if len(parts) != 2 ||
				strings.TrimSpace(parts[0]) == "" ||
				strings.TrimSpace(parts[1]) == "" {

				return config, errors.New(
					"--name must contain two non-empty names separated by a comma",
				)
			}

			config.NameX = strings.TrimSpace(parts[0])
			config.NameO = strings.TrimSpace(parts[1])

		case "--size":
			if i+1 >= len(args) {
				return config, errors.New("missing value for --size")
			}

			i++
			sizeWasSet = true

			size, err := strconv.Atoi(args[i])

			if err != nil || size < 3 {
				return config, errors.New("--size must be an integer >= 3")
			}

			config.Size = size

		default:
			return config, errors.New("unknown flag: " + arg)
		}
	}

	if playersMode == aiMode {
		return config, errors.New("choose exactly one of --players or --ai")
	}

	if aiMode && sizeWasSet {
		return config, errors.New(
			"--ai and --size cannot be combined (AI is 3×3 only)",
		)
	}

	if playersMode {
		config.Mode = "players"
	}

	if aiMode {
		config.Mode = "ai"
	}

	return config, nil
}
