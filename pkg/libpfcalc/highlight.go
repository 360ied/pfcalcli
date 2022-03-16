package libpfcalc

import (
	"strconv"
	"strings"
)

func HighlightANSI(text string, functions map[string]Operator) string {
	const numberColor = "\u001b[32m"   // green
	const operatorColor = "\u001b[34m" // blue
	const commandColor = "\u001b[37m"  // white
	const functionColor = "\u001b[36m" // cyan
	const invalidColor = "\u001b[31m"  // red

	w := new(strings.Builder)

	tokens := strings.Split(text, " ")

	prevTokenWasFunc := false

	for _, v := range tokens {
		if prevTokenWasFunc {
			w.WriteString(functionColor)
			w.WriteString(v)
			prevTokenWasFunc = false
		} else if _, err := strconv.ParseFloat(v, 64); err == nil {
			// token is a number
			w.WriteString(numberColor)
			w.WriteString(v)
		} else if _, in := operators[v]; in {
			// token is an operator
			w.WriteString(operatorColor)
			w.WriteString(v)
		} else if _, in := functions[v]; in {
			// token is a function
			w.WriteString(functionColor)
			w.WriteString(v)
		} else if v == "#func" {
			// token is a func command
			w.WriteString(commandColor)
			w.WriteString(v)
			prevTokenWasFunc = true
		} else {
			// token is invalid
			w.WriteString(invalidColor)
			w.WriteString(v)
		}

		// readd space
		w.WriteString(" ")
	}

	// reset color
	w.WriteString("\u001b[0m")

	return w.String()
}
