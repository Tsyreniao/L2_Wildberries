package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	// Получение точного времени с использованием библиотеки NTP
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		log.Printf("Ошибка при получении времени: %v\n", err)
		os.Exit(1)
	}

	// Получение текущего локального времени
	localTime := time.Now()

	// Вывод
	fmt.Println("Точное время:", ntpTime)
	fmt.Println("Локальное время:", localTime)
}
