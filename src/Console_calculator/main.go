package main

import (
	"fmt"
	"strconv"
	"strings"
)

func InputOperand(side string) float64 {
	for {
		input := ""
		fmt.Println("Input " + side + " operand:")
		fmt.Scan(&input)
		if tmp, err := strconv.ParseFloat(input, 64); err == nil {
			return tmp
		} else {
			fmt.Println("Invalid input")
		}
	}
}
func InputTypeOperation() string {
	for {
		input := ""
		fmt.Println("Input operation")
		fmt.Scan(&input)
		if len(input) > 1 || !(strings.ContainsAny("+-/*", input)) {
			fmt.Println("Invalid input")
		} else {
			return input
		}
	}
}
func Handler(operand_1 float64, operand_2 float64, type_operation string) float64 {
	switch type_operation[0] {
	case '-':
		return operand_1 - operand_2
	case '+':
		return operand_1 + operand_2
	case '*':
		return operand_1 * operand_2
	case '/':
		return operand_1 / operand_2
	default:
		return operand_1 + operand_2
	}
}

func main() {
	operand_1 := InputOperand("left")
	type_operation := InputTypeOperation()
	operand_2 := 0.0
	for {
		operand_2 = InputOperand("right")
		if operand_2 == 0.0 && type_operation[0] == '/' {
			fmt.Println("Invalid input")
			continue
		}
		break
	}

	fmt.Printf("%.3f", Handler(operand_1, operand_2, type_operation))

}
