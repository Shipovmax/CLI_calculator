package main

import (
	"fmt"
	"os"
	"strconv"
)

func parseArgs(args []string) (float64, string, float64, error) {
	if len(args) < 3 {
		return 0, "", 0, fmt.Errorf("need 3 arguments: number operator number")
	}

	number1, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid first number: %s", args[0])
	}

	sign := args[1]

	number2, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid second number: %s", args[2])
	}

	return number1, sign, number2, nil
}

func calculate(a float64, op string, b float64) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", op)
	}
}

func main() {
	a, op, b, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	result, err := calculate(a, op, b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
