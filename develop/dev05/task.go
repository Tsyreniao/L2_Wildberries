package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Значения ключей
var (
	fAfter      int
	fBefore     int
	fContext    int
	fCount      bool
	fIgnoreCase bool
	fInvert     bool
	fFixed      bool
	fLineNum    bool
)

// Считывание ключей при запуске программы
func init() {
	flag.IntVar(&fAfter, "A", 0, "print +N lines after")
	flag.IntVar(&fBefore, "B", 0, "print +N lines before")
	flag.IntVar(&fContext, "C", 0, "print +N lines before and after")
	flag.BoolVar(&fCount, "c", false, "print lines number")
	flag.BoolVar(&fIgnoreCase, "i", false, "ignore case")
	flag.BoolVar(&fInvert, "v", false, "removing insted of match")
	flag.BoolVar(&fFixed, "F", false, "exact match with the line")
	flag.BoolVar(&fLineNum, "n", false, "print line number")

}

func myGrep(pattern string) {
	// Слова в строке
	var patterns []string

	// Если нет флага fFixed, то разбиваем строку на слова
	if !fFixed {
		patterns = strings.Fields(pattern)
	}

	// Открываем файл для чтения (или используем стандартный ввод, если файл не указан)
	var file *os.File
	if flag.NArg() > 1 {
		var err error
		file, err = os.Open(flag.Arg(1))
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()
	} else {
		file = os.Stdin
	}

	// Если флаг FignoreCase установлен, то переводим паттерн в нижний регистр
	if fIgnoreCase {
		pattern = strings.ToLower(pattern)
	}

	// Заполняем значения fBefore и fAfter
	if fContext > 0 {
		if fContext > fAfter {
			fAfter = fContext
		}
		if fContext > fBefore {
			fBefore = fContext
		}
	}

	// Инициализируем счетчик совпадений
	matchCount := 0

	// Инициализируем номер последней строки с совпадением
	lastMatchLine := -1

	// Инициализируем номер последней напечатанной строки без совпадения
	lastNotMathchLine := 0

	// Объявляем переменную для сохранения строк до совпадения
	beforeLines := make(map[int]string, 0)

	// Номер текущей строки
	lineNumVal := 1

	// Создаем сканер для чтения файла построчно
	scanner := bufio.NewScanner(file)

	// Проходим по строкам файла и выполняем фильтрацию
	for scanner.Scan() {
		rawLine := scanner.Text()
		var line string

		// Если флаг FignoreCase установлен, то переводим строку в нижний регистр
		if fIgnoreCase {
			line = strings.ToLower(rawLine)
		} else {
			line = rawLine
		}

		// Проверяем совпадение в строке в зависимости от флагов
		match := false
		if fFixed {
			// Точное совпадение со строкой
			if strings.Contains(line, pattern) {
				match = true
			}
		} else {
			// Совпадение хотябы одного слова из строки
			for _, word := range patterns {
				if strings.Contains(line, word) {
					match = true
					break
				}
			}
		}

		// Если флаг fInvert установлен, инвертируем совпадение
		if fInvert {
			match = !match
		}

		// Удаляем не нужные строки до совпадения
		if len(beforeLines) > fBefore {
			delete(beforeLines, lineNumVal-fBefore-1)
		}

		if match {
			// Увеличиваем счетчик совпадений
			matchCount++

			// Сохраняем номер строки с совпадением
			lastMatchLine = lineNumVal

			// Печатаем строки до совпадения
			for i := lineNumVal - fBefore; i < lineNumVal; i++ {
				if _, ok := beforeLines[i]; ok && i > lastNotMathchLine {
					// Печатаем номер строки, если установлен флаг fLineNum
					if fLineNum {
						fmt.Printf("  %d:", i)
					}
					// Печатаем строку
					fmt.Println(beforeLines[i])
				}
			}

			// Печатаем номер строки, если установлен флаг fLineNum
			if fLineNum {
				fmt.Printf("->%d:", lineNumVal)
			}
			// Печатаем строку
			fmt.Println(rawLine)

			for k := range beforeLines {
				delete(beforeLines, k)
			}
		} else {
			// Сохраняем строки до совпадения
			beforeLines[lineNumVal] = rawLine

			if matchCount > 0 && lineNumVal-lastMatchLine <= fAfter {
				// Печатаем номер строки, если установлен флаг fLineNum
				if fLineNum {
					fmt.Printf("  %d:", lineNumVal)
				}
				// Печатаем строку
				fmt.Println(beforeLines[lineNumVal])

				// Сохраняем номер последней напечатанной строки без совпадения
				lastNotMathchLine = lineNumVal
			}
		}

		lineNumVal++
	}

	// Проверяем ошибки сканера
	if err := scanner.Err(); err != nil && err != io.EOF {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	// Выводим количество совпадений, если установлен флаг fCount
	if fCount {
		fmt.Println("Match count:", matchCount)
	}
}

func main() {
	// Определяем флаги
	flag.Parse()

	// Получаем паттерн для поиска (первый непосредственный аргумент после флагов)
	pattern := flag.Arg(0)

	// Проверяем, что паттерн задан
	if pattern == "" {
		fmt.Println("No pattern specified")
		os.Exit(1)
	}

	// Запускаем функцию
	myGrep(pattern)
}
