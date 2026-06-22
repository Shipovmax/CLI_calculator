package main

import (
	"fmt"
	"os"
	"strconv"
)

func parseArgs(args []string) (float64, string, float64, error) {
	if len(args) != 3 {
		return 0, "", 0, fmt.Errorf("нужно ровно 3 аргумента, получено: %d", len(args))
	}

	number1, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return 0, "", 0, fmt.Errorf("неверный первый номер: %s", args[0])
	}

	op := args[1]

	number2, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return 0, "", 0, fmt.Errorf("неверный второй номер: %s", args[2])
	}

	return number1, op, number2, nil
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
			return 0, fmt.Errorf("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("неизвестный оператор: %s", op)
	}
}

func main() {
	a, op, b, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ошибка: %s\n", err)
		os.Exit(1)
	}

	result, err := calculate(a, op, b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ошибка: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
