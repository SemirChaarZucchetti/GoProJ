package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter the first number: ")
	_, err1 := fmt.Scan(&num1)
	if err1 != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	fmt.Print("Enter the operator (+, -, *, /): ")
	_, _ = fmt.Scan(&operator)

	fmt.Print("Enter the second number: ")
	_, err2 := fmt.Scan(&num2)
	if err2 != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	calculator(num1, num2, operator)

}

func calculator(num1, num2 float64, operator string) (float64, error) {
	var result float64

	switch operator {
	case "+":
		result = add(num1, num2)
		formattedResult := fmt.Sprintf("%.2f", result)
		fmt.Printf("Result: %s\n", formattedResult)
	case "-":
		result = subtract(num1, num2)
		formattedResult := fmt.Sprintf("%.2f", result)
		fmt.Printf("Result: %s\n", formattedResult)
	case "*":
		result = multiply(num1, num2)
		formattedResult := fmt.Sprintf("%.2f", result)
		fmt.Printf("Result: %s\n", formattedResult)
	case "/":
		result = divide(num1, num2)
		formattedResult := fmt.Sprintf("%.2f", result)
		fmt.Printf("Result: %s\n", formattedResult)
	default:
		fmt.Println("Invalid operator. Please use +, -, *, or /.")
		return 0, fmt.Errorf("invalid operator")
	}

	return result, nil
}

func add(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) float64 {
	if y == 0 {
		fmt.Println("Cannot divide by zero.")
		return 0

	}
	return x / y
}
