package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToArabic(roman string) (int, error) {
	romanNumerals := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	val, exists := romanNumerals[roman]
	if !exists {
		return 0, fmt.Errorf("неверное римское число: %s", roman)
	}
	return val, nil
}

func isRomanNumber(num string) bool {
	romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, r := range romanNumerals {
		if r == num {
			return true
		}
	}
	return false
}

func eval(expr string) (int, error) {
	operators := []string{"+", "-", "*", "/"}
	var operator string
	for _, op := range operators {
		if strings.Contains(expr, op) {
			operator = op
			break
		}
	}

	if operator == "" {
		return 0, fmt.Errorf("оператор не найден в выражении")
	}

	parts := strings.Split(expr, operator)
	if len(parts) != 2 {
		return 0, fmt.Errorf("неверный формат выражения")
	}

	operand1, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil || operand1 < 1 || operand1 > 10 {
		if !isArabicNumber(strings.TrimSpace(parts[0])) {
			romanVal, err := romanToArabic(strings.TrimSpace(parts[0]))
			if err != nil {
				return 0, fmt.Errorf("неверный первый операнд: %s", parts[0])
			}
			operand1 = romanVal
		} else {
			return 0, fmt.Errorf("неверный первый операнд: %s", parts[0])
		}
	}

	operand2, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil || operand2 < 1 || operand2 > 10 {
		if !isArabicNumber(strings.TrimSpace(parts[1])) {
			romanVal, err := romanToArabic(strings.TrimSpace(parts[1]))
			if err != nil {
				return 0, fmt.Errorf("неверный второй операнд: %s", parts[1])
			}
			operand2 = romanVal
		} else {
			return 0, fmt.Errorf("неверный второй операнд: %s", parts[1])
		}
	}

	switch operator {
	case "+":
		return operand1 + operand2, nil
	case "-":
		result := operand1 - operand2
		if result < 0 {
			return 0, fmt.Errorf("результат операции отрицательный: %d", result)
		}
		return result, nil
	case "*":
		return operand1 * operand2, nil
	case "/":
		if operand2 == 0 {
			return 0, fmt.Errorf("невозможно делить на ноль")
		}
		result := operand1 / operand2
		return result, nil
	default:
		return 0, fmt.Errorf("недопустимый оператор: %s", operator)
	}
}

func isArabicNumber(num string) bool {
	arabicNumbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for _, arabic := range arabicNumbers {
		if arabic == num {
			return true
		}
	}
	return false
}

func main() {
	fmt.Print("Введите выражение: ")
	var expr string
	fmt.Scanln(&expr)

	result, err := eval(expr)
	if err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}

	fmt.Println("Результат:", result)
}
