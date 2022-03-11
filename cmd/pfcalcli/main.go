package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"pfcalcli/internal/intmath"
	"pfcalcli/internal/stackutil"
	"pfcalcli/pkg/libpfcalc"

	"golang.org/x/term"
)

var build = "DEVELOPMENT_BUILD"

func moveLeft() {
	_, _ = os.Stdout.WriteString("\u001b[1000D") // move all the way left
}

func moveCursor(pos int) {
	_, _ = os.Stdout.WriteString("\u001b[" + strconv.Itoa(pos) + "C")
}

func main() {
	_, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	_, _ = os.Stdout.WriteString("pfcalcli build " + build + "\n")
	moveLeft()

	var (
		stack        []float64
		history      []string
		historyIndex int
	)

	r := bufio.NewReader(os.Stdin)

	for {
		prompt := strconv.Itoa(len(stack)) + "> "
		_, _ = os.Stdout.WriteString(prompt)

		input := ""
		index := 0
		// read input
		for {
			c, err := r.ReadByte()
			if err != nil {
				panic(err)
			}

			if c == 3 {
				// ctrl+c
				moveLeft()
				_, _ = os.Stdout.WriteString("\nExiting...\n")
				moveLeft()
				os.Exit(0)
			} else if c >= 32 && c <= 126 {
				// printable character
				input = input[:index] + string(c) + input[index:]
				index++
			} else if c == 10 || c == 13 { // CRLF
				history = append(history, input)
				historyIndex = len(history) - 1
				break
			} else if c == 27 { // ESC
				next1, err := r.ReadByte()
				if err != nil {
					panic(err)
				}

				next2, err := r.ReadByte()
				if err != nil {
					panic(err)
				}

				if next1 == 91 {
					switch next2 {
					case 'A': // up
						if len(history) == 0 {
							break
						}
						historyIndex = intmath.Max(0, historyIndex-1)
						input = history[historyIndex]
					case 'B': // down
						if len(history) == 0 {
							break
						}
						if historyIndex == len(history) {
							break
						}
						historyIndex = intmath.Min(len(history)-1, historyIndex+1)
						input = history[historyIndex]
					case 'C': // right
						index = intmath.Min(len(input), index+1)
					case 'D': // left
						index = intmath.Max(0, index-1)
					}
				}
			} else if c == 127 { // backspace
				if index > 0 {
					input = input[:index-1] + input[index:]
					index--
				}
			}

			// TODO: make this not call write so many times
			moveLeft()
			_, _ = os.Stdout.WriteString("\u001b[0K")
			_, _ = os.Stdout.WriteString(prompt + input)
			moveLeft()

			if index > 0 {
				moveCursor(index + len(prompt))
			}
		}

		_, _ = os.Stdout.WriteString("\n")
		moveLeft()

		var err error
		stack, err = libpfcalc.Evaluate(stack, input)
		if err != nil {
			_, _ = os.Stdout.WriteString("evaluate: " + err.Error() + "\n")
			moveLeft()
		}

		_, top, found := stackutil.Pop(stack)
		if found {
			// print with trailing zeroes (and decimal point) removed
			_, _ = os.Stdout.WriteString(strings.TrimRight(strings.TrimRight(strconv.FormatFloat(top, 'f', 6, 64), "0"), ".") + "\n")
		} else {
			_, _ = os.Stdout.WriteString("_\n")
		}
		moveLeft()
	}
}
