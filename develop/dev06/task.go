package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Значения ключей
var (
	fFields    string
	fDelimeter string
	fSeparated bool
)

// Считывание ключей при запуске программы
func init() {
	flag.StringVar(&fFields, "f", "", "column for cut")
	flag.StringVar(&fDelimeter, "d", "\t", "change delimeter")
	flag.BoolVar(&fSeparated, "s", false, "print only delimeted lines")
}

func myCut() {
	// Переменная содержащая номера полей для вывода
	var selectedFields []int

	// Если выбранны поля
	if fFields != "" {
		// Получаем выбранные поля в виде массива
		selectedFieldsStrings := strings.Split(fFields, ",")

		// Сохраняем номера полей
		for _, v := range selectedFieldsStrings {
			field, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			selectedFields = append(selectedFields, field)
		}
	}

	// Создаем сканер для чтения из STDIN
	scanner := bufio.NewScanner(os.Stdin)

	// Проходим по строкам файла и выполняем фильтрацию
	for scanner.Scan() {
		line := scanner.Text()

		// Пропускаем итерацию если не найден делиметер и флаг fSeparated поднят
		if fSeparated && !strings.Contains(line, fDelimeter) {
			continue
		}

		// Делим строку с помощью разделителя
		words := strings.Split(line, fDelimeter)

		// Сохраняем номера слов для проверки на существование
		wordsExisting := make(map[int]bool)
		for k := range words {
			wordsExisting[k] = true
		}

		// Если заданы поля для вывода
		if fFields != "" {
			// Выводим все заданные поля если они существуют
			for _, v := range selectedFields {
				if _, ok := wordsExisting[v-1]; ok {
					fmt.Print(words[v-1], " ")
				}
			}
		}

		fmt.Print("\n")
	}
}

func main() {
	// Определяем флаги
	flag.Parse()

	// Запускаем функцию
	myCut()
}
