package libpfcalc

import "pfcalcli/libpfcalc/stackutil"

func opAdd(stack []float64) ([]float64, error) {
	var (
		x, y  float64
		found bool
	)

	stack, x, found = stackutil.Pop(stack)
	if !found {
		return nil, ErrStackUnderflow
	}

	stack, y, found = stackutil.Pop(stack)
	if !found {
		return nil, ErrStackUnderflow
	}

	stack = stackutil.Push(stack, x+y)

	return stack, nil
}

func opSub(stack []float64) ([]float64, error) {
	var (
		x, y  float64
		found bool
	)

	stack, x, found = stackutil.Pop(stack)
	if !found {
		return nil, ErrStackUnderflow
	}

	stack, y, found = stackutil.Pop(stack)
	if !found {
		return nil, ErrStackUnderflow
	}

	stack = stackutil.Push(stack, y-x) // 1 2 - eq 1 - 2

	return stack, nil
}
