package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	// Парсим флаг fTimeout
	fTimeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	// Получаем хоста и порт из командной строки
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: go-telnet [--timeout=<timeout>] <host> <port>")
		return
	}
	host := args[0]
	port := args[1]

	// Создаем канал отслеживающий нажатие Ctrl+C
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Подключаемся к серверу
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", host, port), *fTimeout)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Копирование данных из STDIN на сервер
	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil && err != io.EOF {
			fmt.Println("Error sending data to server:", err)
		}
	}()

	// Копирование данных с сервера
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil && err != io.EOF {
			fmt.Println("Error receiving data from server:", err)
		}
	}()

	// Ожидание нажатия Ctrl+C или завершение fTimeout
	select {
	case <-signalChan:
		conn.Close()
	case <-time.After(*fTimeout):
		conn.Close()
	}
}
