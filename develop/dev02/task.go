package main

import (
	"fmt"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Структура для двух соседних символов
type myChar struct {
	value rune

	isExist  bool
	isNumber bool
}

func myFunc(s string) string {
	var answer string
	ch1 := myChar{
		isExist:  false,
		isNumber: false,
	}
	ch2 := myChar{
		isExist:  false,
		isNumber: false,
	}

	var isBackslash bool = false

	for _, v := range s {
		// Если нашли \, то следующая итерация + запоминаем что предыдущий символ \
		if !isBackslash && v == '\\' {
			isBackslash = true
			continue
		}

		// Если первого символа не сущемтвует, то сохраняем значение
		// Иначе если второго символа не сущемтвует, то сохраняем значение
		if !ch1.isExist {
			ch1.value = v
			ch1.isExist = true

			// Если предыдущий символ \, то символ не может быть числом
			// Иначе помечаем символ как число или не число
			if isBackslash {
				ch1.isNumber = false
				isBackslash = false
			} else if unicode.IsDigit(v) {
				ch1.isNumber = true
			} else {
				ch1.isNumber = false
			}
		} else if !ch2.isExist {
			ch2.value = v
			ch2.isExist = true

			// Если предыдущий символ \, то символ не может быть числом
			// Иначе помечаем символ как число или не число
			if isBackslash {
				ch2.isNumber = false
				isBackslash = false
			} else if unicode.IsDigit(v) {
				ch2.isNumber = true
			} else {
				ch2.isNumber = false
			}
		}

		// Если два символа существуют
		if ch1.isExist && ch2.isExist {
			// Комбинация символ - символ
			if !ch1.isNumber && !ch2.isNumber {
				// Сохраняем 1й символ
				// 2й символ ставим на место 1ого
				// "Удаляем" 2й символ
				answer = answer + string(ch1.value)
				ch1 = ch2
				ch2.isExist = false
			}
			// Комбинация символ - число
			if !ch1.isNumber && ch2.isNumber {
				// Сохраняем 1й символ N раз
				for i := 0; i < int(ch2.value-'0'); i++ {
					answer = answer + string(ch1.value)
				}
				// "Удаляем" 1й и 2й символ
				ch1.isExist = false
				ch2.isExist = false
			}
			// Комбинация число - символ
			if ch1.isNumber && !ch2.isNumber {
				// 2й символ ставим на место 1ого
				// "Удаляем" 2й символ
				ch1 = ch2
				ch2.isExist = false
			}
			// Комбинация число - число
			if ch1.isNumber && ch2.isNumber {
				// "Удаляем" 1й и 2й символ
				ch1.isExist = false
				ch2.isExist = false
			}
		}
	}

	// Если остался символ, то сохраняем его
	if ch1.isExist && !ch1.isNumber {
		answer = answer + string(ch1.value)
	}

	// Вывод результата
	return answer
}

func main() {
	// Строка для распаковки
	var str string = "qwe\\\\55"

	// Запускаем функцию
	newStr := myFunc(str)

	// Вывод
	fmt.Println(str)
	fmt.Println(newStr)
}
