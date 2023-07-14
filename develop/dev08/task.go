package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func myShell() {
	// Создаем сканер для чтения из STDIN
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		// Сохраняем строку и делим её на массив слов
		input := scanner.Text()
		args := strings.Fields(input)

		// Переходим к следующей строке, если найдено 0 слов
		if len(args) == 0 {
			continue
		}

		// Заплняем стркутуру Cmd названием команды и аргументами
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Исполняем команду
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}

func myNetcat() {
	// Отпределяем кол-во введеных аргументов
	if len(os.Args) < 3 {
		fmt.Println("Usage: nc <host> <port>")
		return
	}

	// Сохраняем хост и прот
	host := os.Args[1]
	port := os.Args[2]

	// Подключаемся по аресу host:port по tcp
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Отправляем введенные данные в соединение
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		_, err := fmt.Fprintf(conn, "%s\n", input)
		if err != nil {
			fmt.Println("Error sending data:", err)
			break
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading input:", scanner.Err())
	}
}

func main() {
	//myShell()

	myNetcat()
}
