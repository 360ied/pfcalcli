package libpfcalc

import (
	"pfcalcli/internal/degconv"
	"pfcalcli/internal/stackutil"
)

func mathShim(f func(n float64) float64) func(stack []float64) ([]float64, error) {
	return func(stack []float64) ([]float64, error) {
		var (
			n     float64
			found bool
		)

		stack, n, found = stackutil.Pop(stack)
		if !found {
			return nil, ErrStackUnderflow
		}

		val := f(n)

		stack = stackutil.Push(stack, val)

		return stack, nil
	}
}

func degRadShim(f func(rad float64) float64) func(deg float64) float64 {
	return func(deg float64) float64 {
		var n float64

		n = degconv.DegreesToRadians(deg)
		n = f(n)

		return n
	}
}

func radDegShim(f func(deg float64) float64) func(deg float64) float64 {
	return func(rad float64) float64 {
		var n float64

		n = f(rad)
		n = degconv.RadiansToDegrees(n)

		return n
	}
}

func constantShim(c float64) func(stack []float64) ([]float64, error) {
	return func(stack []float64) ([]float64, error) {
		stack = stackutil.Push(stack, c)

		return stack, nil
	}
}

func combineOps(ops []Operator) func(stack []float64) ([]float64, error) {
	return func(stack []float64) ([]float64, error) {
		newStack, err := evalOps(stack, ops)
		if err != nil {
			return nil, err
		}

		return newStack, nil
	}
}
