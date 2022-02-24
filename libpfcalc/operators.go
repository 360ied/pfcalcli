package libpfcalc

import (
	"fmt"

	"pfcalcli/libpfcalc/stackutil"
)

func opSwap(stack []float64) ([]float64, error) {
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

	stack = stackutil.Push(stack, x, y)

	return stack, nil
}

func opDup(stack []float64) ([]float64, error) {
	var (
		n     float64
		found bool
	)

	_, n, found = stackutil.Pop(stack)
	if !found {
		return nil, ErrStackUnderflow
	}

	stack = stackutil.Push(stack, n)

	return stack, nil
}

func opOver(stack []float64) ([]float64, error) {
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

	stack = stackutil.Push(stack, x, y, x)

	return stack, nil
}

func opRot(stack []float64) ([]float64, error) {
	var (
		x, y, z float64
		found   bool
	)

	stack, x, found = stackutil.Pop(stack)
	if !found {
		return nil, ErrStackUnderflow
	}

	stack, y, found = stackutil.Pop(stack)
	if !found {
		return nil, ErrStackUnderflow
	}

	stack, z, found = stackutil.Pop(stack)
	if !found {
		return nil, ErrStackUnderflow
	}

	stack = stackutil.Push(stack, y, x, z)

	return stack, nil
}

func opDrop(stack []float64) ([]float64, error) {
	var found bool

	stack, _, found = stackutil.Pop(stack)
	if !found {
		return nil, ErrStackUnderflow
	}

	return stack, nil
}

func opPrint(stack []float64) ([]float64, error) {
	fmt.Println(stack)

	return stack, nil
}

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
