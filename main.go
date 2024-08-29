package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Преобразование римских чисел в арабские
func romanToInt(s string) (int, error) {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	prevValue := 0

	// Приводим входное значение к верхнему регистру
	s = strings.ToUpper(s)

	for _, char := range s {
		currentValue, exists := romanNumerals[char]
		if !exists {
			return 0, fmt.Errorf("недопустимый символ: %c", char)
		}

		// Если текущее значение больше предыдущего, значит мы попали в ситуацию, когда
		// нужно вычесть предыдущее значение (например, IV, IX и т.п.)
		if currentValue > prevValue {
			total += currentValue - 2*prevValue // Вычитаем двойное предыдущее значение
		} else {
			total += currentValue
		}
		prevValue = currentValue
	}

	return total, nil
}

func main() {
	input := make([]string, 4)
	fmt.Println("Введите строку в формате 'число операция число':")
	_, err := fmt.Scanln(&input[0], &input[1], &input[2])
	if err != nil {
		fmt.Println("Паника! Должно быть всего 2 аргумента и один знак операции", err)
		return
	}

	// Проверяем количество частей
	if input[3] != "" {
		fmt.Println(`Паника! Необходимо ввести строку в формате 'число операция число'`)
		return
	}
	// Преобразуем первое и третье поле в числа
	firstNumber, err1 := strconv.Atoi(input[0])
	secondNumber, err2 := strconv.Atoi(input[2])
	operand := input[1]
	if err1 != nil && err2 != nil {
		var arabic1, arabic2 int
		var result int
		arabic1, _ = romanToInt(input[0])
		arabic2, _ = romanToInt(input[2])
		result = arabic1
		switch operand {
		case "+":
			result += arabic2
		case "-":
			result -= arabic2
		case "/":
			result /= arabic2
		case "*":
			result *= arabic2
		case "default":
			fmt.Println("Паника! Нет такой операции.")
			return
		}
		if result >= 0 {
			fmt.Println(result)
			return
		} else {
			fmt.Println("Паника! Результат операции с римскими числами должен быть больше 0")
		}
	} else if !(err1 == nil && err2 == nil) {
		fmt.Println("Паника! Нужно указывать числа в одном формате. Оба числа арабские или оба числа римские")
	} else {
		result := firstNumber
		switch operand {
		case "+":
			result += secondNumber
		case "-":
			result -= secondNumber
		case "/":
			result /= secondNumber * 1.0
		case "*":
			result *= secondNumber
		case "default":
			fmt.Println("Паника! Нет такой операции.")
			return
		}
		fmt.Println(result)
	}
}
