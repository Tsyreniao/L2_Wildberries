package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func findAnagramSets(words *[]string) *map[string]*[]string {
	// Мапа с ключом - для поиска анаграмм
	anagramSets := make(map[string]*[]string)
	// Мапа с ключом первое встретившееся в словаре слово из множества
	sortedAnagramSets := make(map[string]*[]string)

	// Приводим слова к нижнему регистру
	for k := range *words {
		(*words)[k] = strings.ToLower((*words)[k])
	}

	for _, word := range *words {
		// Сортируем буквы
		sortedWord := sortString(word)

		// Если множество анаграмм уже существует, добавляем слово в него
		if _, ok := anagramSets[sortedWord]; ok {
			*anagramSets[sortedWord] = append(*anagramSets[sortedWord], word)
		} else {
			// Создаем новое множество анаграмм и добавляем его в мапу
			anagramSets[sortedWord] = &[]string{word}
		}
	}

	for _, v := range anagramSets {
		// Множество из одного элемента не сохраяем
		if len(*v) == 1 {
			continue
		}
		// Сохраняем первое слово
		firstWord := (*v)[0]
		// Сортируем массив
		sort.Strings(*v)
		// Сохраняем сортированный массив
		sortedAnagramSets[firstWord] = v
	}

	return &sortedAnagramSets
}

// Функция соритровки букв в слове в алфавитном порядке
func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	// Словарь
	dictionary := []string{"Пятак", "пятка", "Тяпка", "столик", "Листок", "слиток", "Мама", "амам", "Папа"}

	// Запускаем функцию
	anagramSets := findAnagramSets(&dictionary)

	// Вывод анаграмм
	for k, wordSet := range *anagramSets {
		fmt.Printf("%v: %v\n", k, *wordSet)
	}
}
