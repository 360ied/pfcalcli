package libpfcalc

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"pfcalcli/internal/stackutil"
)

var (
	ErrStackUnderflow  = errors.New("libpfcalc: stack underflow")
	ErrInvalidOperator = errors.New("libpfcalc: invalid operator")
)

var operators = map[string]func(stack []float64) ([]float64, error){
	"+":           opAdd,
	"-":           opSub,
	"*":           opMul,
	"/":           opDiv,
	"^":           opPow,
	"swap":        opSwap,
	"dup":         opDup,
	"over":        opOver,
	"rot":         opRot,
	"drop":        opDrop,
	"clr":         opClr,
	"print":       opPrint,
	"abs":         mathShim(math.Abs),
	"acos":        mathShim(math.Acos),
	"acosh":       mathShim(math.Acosh),
	"asin":        mathShim(math.Asin),
	"asinh":       mathShim(math.Asinh),
	"atan":        mathShim(math.Atan),
	"atanh":       mathShim(math.Atanh),
	"cbrt":        mathShim(math.Cbrt),
	"ceil":        mathShim(math.Ceil),
	"cos":         mathShim(math.Cos),
	"cosh":        mathShim(math.Cosh),
	"erf":         mathShim(math.Erf),
	"erfc":        mathShim(math.Erfc),
	"erfcinv":     mathShim(math.Erfcinv),
	"erfinv":      mathShim(math.Erfinv),
	"exp":         mathShim(math.Exp),
	"exp2":        mathShim(math.Exp2),
	"expm1":       mathShim(math.Expm1),
	"floor":       mathShim(math.Floor),
	"gamma":       mathShim(math.Gamma),
	"j0":          mathShim(math.J0),
	"j1":          mathShim(math.J1),
	"log":         mathShim(math.Log),
	"log10":       mathShim(math.Log10),
	"log1p":       mathShim(math.Log1p),
	"log2":        mathShim(math.Log2),
	"logb":        mathShim(math.Logb),
	"round":       mathShim(math.Round),
	"roundtoeven": mathShim(math.RoundToEven),
	"sin":         mathShim(math.Sin),
	"sinh":        mathShim(math.Sinh),
	"sqrt":        mathShim(math.Sqrt),
	"tan":         mathShim(math.Tan),
	"tanh":        mathShim(math.Tanh),
	"trunc":       mathShim(math.Trunc),
	"y0":          mathShim(math.Y0),
	"y1":          mathShim(math.Y1),
	"e":           constantShim(math.E),
	"pi":          constantShim(math.Pi),
	"phi":         constantShim(math.Phi),
	"sqrt2":       constantShim(math.Sqrt2),
	"sqrte":       constantShim(math.SqrtE),
	"sqrtpi":      constantShim(math.SqrtPi),
	"sqrtphi":     constantShim(math.SqrtPhi),
	"ln2":         constantShim(math.Ln2),
	"log2e":       constantShim(math.Log2E),
	"ln10":        constantShim(math.Ln10),
	"log10e":      constantShim(math.Log10E),
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
