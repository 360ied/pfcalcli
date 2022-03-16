package libpfcalc

import (
	"errors"
	"math"
	"os"
	"strconv"

	"pfcalcli/internal/avg"
	"pfcalcli/internal/stackutil"
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

func opClr([]float64) ([]float64, error) {
	return make([]float64, 0), nil
}

func opPrint(stack []float64) ([]float64, error) {
	_, _ = os.Stdout.WriteString("[")
	for _, v := range stack {
		_, _ = os.Stdout.WriteString(" " + strconv.FormatFloat(v, 'f', -1, 64) + " ")
	}
	_, _ = os.Stdout.WriteString("]\n")

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

func opMul(stack []float64) ([]float64, error) {
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

	stack = stackutil.Push(stack, x*y)

	return stack, nil
}

func opDiv(stack []float64) ([]float64, error) {
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

	stack = stackutil.Push(stack, y/x) // 1 2 / eq 1 / 2

	return stack, nil
}

func opPow(stack []float64) ([]float64, error) {
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

	stack = stackutil.Push(stack, math.Pow(y, x))

	return stack, nil
}

func opAvg(stack []float64) ([]float64, error) {
	const MAX_N = 1024

	var (
		nf    float64 // number of numbers to average
		found bool
	)

	stack, nf, found = stackutil.Pop(stack)
	if !found {
		return nil, ErrStackUnderflow
	}

	n := int(nf)
	if n < 1 {
		return nil, errors.New("libpfcalc: opAvg: number of numbers to average is smaller than 1")
	} else if n > MAX_N {
		return nil, errors.New("libpfcalc: opAvg: number of numbers to average is too large! n > MAX_N (" + strconv.Itoa(MAX_N) + ")")
	}

	nums := make([]float64, n)

	stack, found = stackutil.PopN(stack, nums)
	if !found {
		return nil, ErrStackUnderflow
	}

	stack = stackutil.Push(stack, avg.Avg(nums))

	return stack, nil
}

func opLen(stack []float64) ([]float64, error) {
	stack = stackutil.Push(stack, float64(len(stack)))

	return stack, nil
}

func opEq(stack []float64) ([]float64, error) {
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

	var ret float64

	if x == y {
		ret = 1
	} else {
		ret = 0
	}

	stack = stackutil.Push(stack, ret)

	return stack, nil
}
