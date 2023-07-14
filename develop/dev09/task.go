package main

import (
	"fmt"
	"os"
	"os/exec"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func myWget(url string) {
	cmd := exec.Command("wget", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
	}
}

func main() {
	myWget("google.com")
}
