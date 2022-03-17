package libpfcalc

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"pfcalcli/internal/cmdparser"
	"pfcalcli/internal/stackutil"
)

var (
	ErrStackUnderflow  = errors.New("libpfcalc: stack underflow")
	ErrInvalidOperator = errors.New("libpfcalc: invalid operator")
)

type Operator = func(stack []float64) ([]float64, error)

var operators = map[string]Operator{
	"+":           opAdd,
	"-":           opSub,
	"*":           opMul,
	"/":           opDiv,
	"^":           opPow,
	"pow":         opPow,
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
	"avg":         opAvg,
	"len":         opLen,
	"=":           opEq,
	"eq":          opEq,
	"cmp":         opCmp,
	"noop":        opNoOp,
}

// Evaluate doesn't modify stack, the returned slice is a new allocation
// If an error is returned, the old stack is returned as well
func Evaluate(stack []float64, expressionStr string, functions map[string]Operator) ([]float64, error) {
	cmdName, cmdPos, err := cmdparser.Parse(expressionStr)
	if err != nil {
		if err == cmdparser.ErrNotCommand {
			ops, err := parse(expressionStr, functions)
			if err != nil {
				return stack, err
			}

			stack, err = evalOps(stack, ops)
			return stack, err
		} else {
			return stack, err
		}
	}

	if cmdName == "func" {
		if len(cmdPos) < 2 {
			// function has no body (or name)
			return stack, errors.New("libpfcalc: Evaluate: function has no body (or name)")
		}

		funcName := cmdPos[0]
		funcContent := strings.Join(cmdPos[1:], " ")

		if err := RegisterFunction(functions, funcName, funcContent); err != nil {
			return stack, err
		}

		return stack, nil
	}

	return stack, errors.New("libpfcalc: Evaluate: unknown function")
}

func RegisterFunction(functions map[string]Operator, name, expressionStr string) error {
	ops, err := parse(expressionStr, functions)
	if err != nil {
		return err
	}

	fu := combineOps(ops)

	functions[name] = fu

	return nil
}

func parse(s string, functions map[string]Operator) ([]Operator, error) {
	var ops []Operator

	for _, tok := range strings.Split(s, " ") {
		tok = strings.TrimSpace(tok)
		if len(tok) == 0 {
			// skip empty tokens
			continue
		}

		n, err := strconv.ParseFloat(tok, 64)
		if err != nil {
			// value is not a float
			op, in := operators[tok]
			if !in {
				// value is not an operator
				fu, in := functions[tok]
				if !in {
					// value is not a function
					return nil, ErrInvalidOperator
				} else {
					// value is a function
					ops = append(ops, fu)
				}
			} else {
				// value is an operator
				ops = append(ops, op)
			}
		} else {
			// value is a float
			ops = append(ops, constantShim(n))
		}
	}

	return ops, nil
}

func evalOps(stack []float64, ops []Operator) ([]float64, error) {
	newStack := stackutil.Clone(stack)

	for _, op := range ops {
		var err error
		newStack, err = op(newStack)
		if err != nil {
			return stack, err
		}
	}

	return newStack, nil
}
