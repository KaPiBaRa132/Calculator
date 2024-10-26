package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Calc - функция вычисления арифметического выражения
func Calc(expression string) (float64, error) {
	// Удаление пробелов
	expression = strings.ReplaceAll(expression, " ", "")

	// Проверка на корректность записи выражения
	if !isValidExpression(expression) {
		return 0, errors.New("invalid expression")
	}

	// Вычисление выражения в скобках
	expression = processParentheses(expression)

	// Вычисление умножения и деления
	expression = processMultiplicationAndDivision(expression)

	// Вычисление сложения и вычитания
	expression = processAdditionAndSubtraction(expression)

	// Преобразование результата в float64
	result, err := strconv.ParseFloat(expression, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}

// isValidExpression - функция проверки корректности записи выражения
func isValidExpression(expression string) bool {
	// Проверка на наличие скобок
	if strings.Count(expression, "(") != strings.Count(expression, ")") {
		return false
	}

	// Проверка на наличие недопустимых символов
	for _, char := range expression {
		if !((char >= '0' && char <= '9') ||
			(char == '+' || char == '-' || char == '*' || char == '/' || char == '(' || char == ')')) {
			return false
		}
	}

	return true
}

// processParentheses - функция вычисления выражения в скобках
func processParentheses(expression string) string {
	for i := 0; i < len(expression); i++ {
		if expression[i] == '(' {
			j := i + 1
			count := 1
			for ; j < len(expression) && count > 0; j++ {
				if expression[j] == '(' {
					count++
				} else if expression[j] == ')' {
					count--
				}
			}
			if count == 0 {
				subExpression := expression[i+1 : j-1]
				result, err := Calc(subExpression)
				if err != nil {
					return "" // или return err, если нужно вернуть ошибку
				}
				// Используем strconv.FormatFloat для корректного преобразования в строку
				expression = expression[:i] + strconv.FormatFloat(result, 'f', -1, 64) + expression[j:]
			}
		}
	}
	return expression
}

// processMultiplicationAndDivision - функция вычисления умножения и деления
func processMultiplicationAndDivision(expression string) string {
	for i := 0; i < len(expression); i++ {
		if expression[i] == '*' || expression[i] == '/' {
			leftOperand := expression[i-1]
			rightOperand := expression[i+1]
			result := calculate(leftOperand, expression[i], rightOperand)
			// Используем strconv.FormatFloat для корректного преобразования в строку
			expression = expression[:i-1] + strconv.FormatFloat(result, 'f', -1, 64) + expression[i+2:]
			i -= 2
		}
	}
	return expression
}

// processAdditionAndSubtraction - функция вычисления сложения и вычитания
func processAdditionAndSubtraction(expression string) string {
	for i := 0; i < len(expression); i++ {
		if expression[i] == '+' || expression[i] == '-' {
			leftOperand := expression[i-1]
			rightOperand := expression[i+1]
			result := calculate(leftOperand, expression[i], rightOperand)
			// Используем strconv.FormatFloat для корректного преобразования в строку
			expression = expression[:i-1] + strconv.FormatFloat(result, 'f', -1, 64) + expression[i+2:]
			i -= 2
		}
	}
	return expression
}

// calculate - функция вычисления арифметической операции
func calculate(leftOperand, operator, rightOperand byte) float64 {
	l, _ := strconv.ParseFloat(string(leftOperand), 64)
	r, _ := strconv.ParseFloat(string(rightOperand), 64)

	switch operator {
	case '*':
		return l * r
	case '/':
		if r == 0 {
			return 0 // или return err, если нужно вернуть ошибку
		}
		return l / r
	case '+':
		return l + r
	case '-':
		return l - r
	}
	return 0
}

func main() {
	expression := "(1+2)*3/4"
	result, err := Calc(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}
