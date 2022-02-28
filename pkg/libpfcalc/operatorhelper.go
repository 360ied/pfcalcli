package libpfcalc

import "pfcalcli/internal/stackutil"

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

func constantShim(c float64) func(stack []float64) ([]float64, error) {
	return func(stack []float64) ([]float64, error) {
		stack = stackutil.Push(stack, c)

		return stack, nil
	}
}
