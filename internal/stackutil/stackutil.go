package stackutil

func Push(stack []float64, value ...float64) []float64 {
	stack = append(stack, value...)
	return stack
}

// bool is false if the stack is empty
func Pop(stack []float64) ([]float64, float64, bool) {
	if len(stack) == 0 {
		return stack, 0, false
	}

	value := stack[len(stack)-1]

	return stack[0 : len(stack)-1], value, true
}

func Clone(stack []float64) []float64 {
	newStack := make([]float64, len(stack))
	copy(newStack, stack)
	return newStack
}
