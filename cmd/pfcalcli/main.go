package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"

	"pfcalcli/internal/intmath"
	"pfcalcli/internal/stackutil"
	"pfcalcli/pkg/libpfcalc"

	"golang.org/x/term"
)

var build = "DEVELOPMENT_BUILD"

func moveLeft(w io.StringWriter) error {
	_, err := w.WriteString("\u001b[1000D") // move all the way left
	return err
}

func moveCursor(w io.StringWriter, pos int) error {
	_, err := w.WriteString("\u001b[" + strconv.Itoa(pos) + "C")
	return err
}

func promptStr(stackLen int, normalMode bool) string {
	prompt := strconv.Itoa(stackLen)
	if normalMode {
		prompt += "N "
	} else {
		prompt += "> "
	}

	return prompt
}

func main() {
	w := bufio.NewWriter(os.Stdout)

	_, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	if _, err := w.WriteString("pfcalcli build " + build + "\n"); err != nil {
		panic(err)
	}
	if err := moveLeft(w); err != nil {
		panic(err)
	}
	if err := w.Flush(); err != nil {
		panic(err)
	}

	var (
		stack        []float64
		history      []string
		historyIndex int
		normalMode   bool

		functions = map[string]libpfcalc.Operator{}
	)

	r := bufio.NewReader(os.Stdin)

	for {
		prompt := promptStr(len(stack), normalMode)
		_, err := w.WriteString(prompt)
		if err != nil {
			panic(err)
		}
		if err := w.Flush(); err != nil {
			panic(err)
		}

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
				if err := moveLeft(w); err != nil {
					panic(err)
				}
				if _, err := w.WriteString("\nExiting...\n"); err != nil {
					panic(err)
				}
				if err := moveLeft(w); err != nil {
					panic(err)
				}
				if err := w.Flush(); err != nil {
					panic(err)
				}
				os.Exit(0)
			} else if c == 9 {
				// tab
				normalMode = !normalMode
				prompt = promptStr(len(stack), normalMode)
			} else if c >= 32 && c <= 126 {
				// printable character
				if normalMode {
					if c == '$' {
						// move cursor to the end of line
						index = len(input)
					} else if c == '^' {
						// move cursor to the start of line
						index = 0
					}
				} else {
					// add char to input
					input = input[:index] + string(c) + input[index:]
					index++
				}
			} else if c == 10 || c == 13 { // CRLF
				history = append(history, input)
				historyIndex = len(history)
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
						index = len(input)
					case 'B': // down
						if len(history) == 0 {
							break
						}
						if historyIndex == len(history) {
							break
						}
						historyIndex = intmath.Min(len(history)-1, historyIndex+1)
						input = history[historyIndex]
						index = len(input)
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

			if err := moveLeft(w); err != nil {
				panic(err)
			}
			if _, err := w.WriteString("\u001b[0K" + prompt + libpfcalc.HighlightANSI(input, functions)); err != nil {
				panic(err)
			}
			if err := moveLeft(w); err != nil {
				panic(err)
			}
			if err := moveCursor(w, index+len(prompt)); err != nil {
				panic(err)
			}
			if err := w.Flush(); err != nil {
				panic(err)
			}
		}

		if _, err := w.WriteString("\n"); err != nil {
			panic(err)
		}
		if err := moveLeft(w); err != nil {
			panic(err)
		}
		if err := w.Flush(); err != nil {
			panic(err)
		}

		var err_ error
		stack, err_ = libpfcalc.Evaluate(stack, input, functions)
		if err_ != nil {
			if _, err := w.WriteString("evaluate: " + err_.Error() + "\n"); err != nil {
				panic(err)
			}
			if err := moveLeft(w); err != nil {
				panic(err)
			}
			if err := w.Flush(); err != nil {
				panic(err)
			}
		}

		_, top, found := stackutil.Pop(stack)
		if found {
			// print with trailing zeroes (and decimal point) removed
			if _, err :=
				w.WriteString(
					strings.TrimRight(
						strings.TrimRight(
							strconv.FormatFloat(top, 'f', 6, 64),
							"0"),
						".") +
						"\n"); err != nil {
				panic(err)
			}
		} else {
			if _, err := w.WriteString("_\n"); err != nil {
				panic(err)
			}
		}
		if err := moveLeft(w); err != nil {
			panic(err)
		}
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}
}
