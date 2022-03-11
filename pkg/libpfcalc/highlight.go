package libpfcalc

import (
	"strconv"
	"strings"
)

func HighlightANSI(text string) string {
	const numberColor = "\u001b[32m"   // green
	const operatorColor = "\u001b[34m" // blue
	const invalidColor = "\u001b[31m"  // red

	w := new(strings.Builder)

	tokens := strings.Split(text, " ")

	for _, v := range tokens {
		if _, err := strconv.ParseFloat(v, 64); err == nil {
			// token is a number
			w.WriteString(numberColor)
			w.WriteString(v)
		} else if _, in := operators[v]; in {
			// token is an operator
			w.WriteString(operatorColor)
			w.WriteString(v)
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
