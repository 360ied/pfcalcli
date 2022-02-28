package main

import (
	"bufio"
	"fmt"
	"os"

	"pfcalcli/libpfcalc"
	"pfcalcli/libpfcalc/stackutil"
)

var build = "2"

func main() {
	fmt.Printf("pfcalcli build %s\n", build)

	var stack []float64

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("%d> ", len(stack))

		if !scanner.Scan() {
			break
		}

		expressionStr := scanner.Text()

		var err error
		stack, err = libpfcalc.Evaluate(stack, expressionStr)
		if err != nil {
			fmt.Printf("evaluate: %v\n", err)
		}

		_, top, found := stackutil.Pop(stack)
		if found {
			fmt.Printf("%.6g\n", top)
		} else {
			fmt.Print("_\n")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("main: scanner error: %v\n", err)
		os.Exit(1)
	}
}
