package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"pfcalcli/internal/stackutil"
	"pfcalcli/pkg/libpfcalc"
)

var build = "DEVELOPMENT_BUILD"

func main() {
	_, _ = os.Stdout.WriteString("pfcalcli build " + build + "\n")

	var stack []float64

	scanner := bufio.NewScanner(os.Stdin)

	for {
		_, _ = os.Stdout.WriteString(strconv.Itoa(len(stack)) + "> ")

		if !scanner.Scan() {
			break
		}

		expressionStr := scanner.Text()

		var err error
		stack, err = libpfcalc.Evaluate(stack, expressionStr)
		if err != nil {
			_, _ = os.Stdout.WriteString("evaluate: " + err.Error() + "\n")
		}

		_, top, found := stackutil.Pop(stack)
		if found {
			// print with trailing zeroes (and decimal point) removed
			_, _ = os.Stdout.WriteString(strings.TrimRight(strings.TrimRight(strconv.FormatFloat(top, 'f', 6, 64), "0"), ".") + "\n")
		} else {
			_, _ = os.Stdout.WriteString("_\n")
		}
	}

	if err := scanner.Err(); err != nil {
		_, _ = os.Stdout.WriteString("main: scanner error: " + err.Error() + "\n")
		os.Exit(1)
	}
}
