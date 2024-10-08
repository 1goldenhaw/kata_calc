/*
Требования:
1) Калькулятор умеет выполнять операции сложения, вычитания, умножения и деления с двумя числами: a + b, a - b, a * b, a / b. Данные передаются в одну строку. Решения, в которых каждое число и арифметическая операция передаются с новой строки, считаются неверными.
2) Калькулятор умеет работать как с арабскими (1, 2, 3, 4, 5…), так и с римскими (I, II, III, IV, V…) числами.
3) Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более. На выходе числа не ограничиваются по величине и могут быть любыми.
4) Калькулятор умеет работать только с целыми числами.
5) Калькулятор умеет работать только с арабскими или римскими цифрами одновременно, при вводе пользователем строки вроде 3 + II калькулятор должен выдать панику и прекратить работу.
6) При вводе римских чисел ответ должен быть выведен римскими цифрами, соответственно, при вводе арабских — ответ ожидается арабскими.
7) При вводе пользователем не подходящих чисел приложение выдаёт панику и завершает работу.
8) При вводе пользователем строки, не соответствующей одной из вышеописанных арифметических операций, приложение выдаёт панику и завершает работу.
9) Результатом операции деления является целое число, остаток отбрасывается.
10) Результатом работы калькулятора с арабскими числами могут быть отрицательные числа и ноль. Результатом работы калькулятора с римскими числами могут быть только положительные числа, если результат работы меньше единицы, программа должна выдать панику.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Преобразование римских чисел в арабские
func romanToInt(s string) (int, error) {
	romanNumerals := map[rune]int{ // Создаём словарь римских чисел с отображением в арабские
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

	s = strings.ToUpper(s) // Для удобства приводим входное значение к верхнему регистру
	_, err3 := checkRoman(s)
	if err3 != nil {
		return 0, fmt.Errorf("неверная запись римского числа")
	}

	for _, char := range s {
		currentValue, exists := romanNumerals[char]
		if !exists { // Возвращаем ошибку, преобразование не удалось - левый символ!
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

// Преобразование арабских чисел в римские
func arabicToRoman(num int) string {
	// массивы значений и римских символов
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""

	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			roman += symbols[i]
			num -= values[i]
		}
	}
	return roman
}

func checkRoman(w string) (int, error) {
	romans := make(map[string]int) // Создаём словарь римских чисел меньше 10
	romans["I"] = 1
	romans["II"] = 2
	romans["III"] = 3
	romans["IV"] = 4
	romans["V"] = 5
	romans["VI"] = 6
	romans["VII"] = 7
	romans["VIII"] = 8
	romans["IX"] = 9
	romans["X"] = 10
	if _, found := romans[w]; found { // Проверяем пришедшее число, есть ли оно в списке
		return 0, nil
	} else {
		return 1, fmt.Errorf("римское число %q не существует", w)
	}
}

func main() {
	fmt.Println("Введите строку в формате 'число операция число':")
	var line string
	scanner := bufio.NewScanner(os.Stdin) // Инициализируем сканер
	if scanner.Scan() {                   // Получаем введенную строку с пробелами
		line = scanner.Text()
	}
	if err := scanner.Err(); err != nil { // Обрабатываем ошибки ввода
		panic("Ошибка при чтении!")
	}
	elements := strings.Fields(line) // Создаём массив слов из строки
	if len(elements) > 4 {           // Если хотя бы на один элемент напечатано больше, чем нам надо - выдаём панику
		panic("Введено 2-х больше аргументов / больше одной операции!")
	} else if len(elements) < 3 {
		panic("Выдача паники, так как строка не является математической операцией (или забыты пробелы).")
	}

	firstNumber, err1 := strconv.Atoi(elements[0]) // Преобразуем первое и третье слова в числа
	secondNumber, err2 := strconv.Atoi(elements[2])
	operand := elements[1]          // Запоминаем в отдельную переменную строку знак операции
	if err1 != nil && err2 != nil { // Если обе строки не удалось преобразовать в число, то проверяем их на Римскость
		var arabic1, arabic2 int
		var result int
		arabic1, err3 := romanToInt(elements[0])
		arabic2, err4 := romanToInt(elements[2])
		if arabic1 > 10 || arabic2 > 10 || arabic1 < 1 || arabic2 < 1 { // Тут проверяем больше ли хотя бы одно из введённых чисел 10 или нет
			panic("Введено число больше 10 или меньше 1.")
		}
		if err3 != nil || err4 != nil { // Тут после проверки на Римскость получаем результат. Если ошибка есть - паника
			panic("Паника! Одно из введённых римских чисел не существует или оно больше 10")
		}
		result = arabic1 // В результат промежуточно помещаем первое число.
		switch operand { // Расшифровываем "операцию" и выполняем её!
		case "+":
			result += arabic2
		case "-":
			result -= arabic2
		case "/":
			result /= arabic2
		case "*":
			result *= arabic2
		case "default": // Если введённый символ операции не распознан - паника!
			panic("Паника! Нет такой операции.")
		}
		if result > 0 { // Перед печатью проверяем, чтобы результат операции с римскими числами был положительным
			var resultRoman string
			resultRoman = arabicToRoman(result) /// Преобразование обратно в римские!
			fmt.Println(resultRoman)
			return
		} else {
			panic("Паника! Результат операции с римскими числами должен быть больше 0")
		}
	} else if !(err1 == nil && err2 == nil) { // Если в число преобразовать удалось только один аргумент
		panic("Паника! Нужно указывать числа в одном формате. Оба числа арабские или оба числа римские")
	} else { // Если в число преобразовать удалось оба слова
		if firstNumber > 10 || firstNumber < 1 || secondNumber > 10 || secondNumber < 1 {
			panic("Введено число больше 10 или меньше 1")
		}
		result := firstNumber
		switch operand {
		case "+":
			result += secondNumber
		case "-":
			result -= secondNumber
		case "/":
			result /= secondNumber
		case "*":
			result *= secondNumber
		case "default":
			panic("Паника! Нет такой операции.")
		}
		fmt.Println(result)
	}
}
