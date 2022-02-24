package libpfcalc

import (
	"errors"
	"strconv"
	"strings"

	"pfcalcli/libpfcalc/stackutil"
)

var (
	ErrStackUnderflow  = errors.New("libpfcalc: stack underflow")
	ErrInvalidOperator = errors.New("libpfcalc: invalid operator")
)

var operators map[string]func(stack []float64) ([]float64, error) = map[string]func(stack []float64) ([]float64, error){
	"+":     opAdd,
	"-":     opSub,
	"*":     opMul,
	"/":     opDiv,
	"swap":  opSwap,
	"dup":   opDup,
	"over":  opOver,
	"rot":   opRot,
	"drop":  opDrop,
	"print": opPrint,
}

// Evaluate doesn't modify stack, the returned slice is a new allocation
// If an error is returned, the old stack is returned as well
func Evaluate(stack []float64, expressionStr string) ([]float64, error) {
	newStack := stackutil.Clone(stack)

	expressions := strings.Split(expressionStr, " ")

	for _, expression := range expressions {
		expression = strings.TrimSpace(expression)

		if len(expression) == 0 {
			continue
		}

		value, err := strconv.ParseFloat(expression, 64)
		if err != nil {
			// value is not a float
			operator, in := operators[expression]
			if !in {
				return stack, ErrInvalidOperator
			}

			newStack, err = operator(newStack)
			if err != nil {
				// operator had an error
				return stack, err
			}

			continue
		}

		newStack = stackutil.Push(newStack, value)
	}

	return newStack, nil
}
