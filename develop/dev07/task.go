package main

import (
	"fmt"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

// Функция создает и возвращает канал done и закрывает его если один из полученных каналов завершился
func or(channels ...<-chan interface{}) <-chan interface{} {
	done := make(chan interface{})

	// Проходим по всем каналам и запускаем горутину отслеживающую закрытие канала
	for k := range channels {
		go func(ch <-chan interface{}) {
			select {
			// Если канал done закрыт, то выходим
			case <-done:
				return
			// Если канал закрыт, то закрываем канал done и выходим
			case <-ch:
				close(done)
				return
			}
		}(channels[k])
	}
	return done
}

// Функция создает и возвращает канал, который закроется через установленное время
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	// Сохраняем время начала работы программы
	start := time.Now()

	// Ожидаем закрытия канала done
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Hour),
		sig(1*time.Minute),
		sig(2*time.Second),
		sig(1*time.Second),
	)

	// Выводим время работы программы
	fmt.Printf("Done after %v", time.Since(start))
}
