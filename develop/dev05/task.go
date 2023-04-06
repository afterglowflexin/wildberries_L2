package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
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

// use case: go run task.go flags source substring

type Config struct {
	A int
	B int
	C int
	c bool
	i bool
	v bool
	f bool
	n bool
}

type App struct {
	source    *os.File
	substring string
	config    Config
}

func NewApp() *App {
	return &App{config: Config{}}
}

func (a *App) Init() {
	source := os.Args[len(os.Args)-2]
	a.substring = os.Args[len(os.Args)-1]

	file, err := os.Open(source)

	if err != nil {
		log.Println("Unable to open the file")
		os.Exit(1)
	}

	a.source = file
}

func main() {

	app := NewApp()
	app.Init()

	output, err := app.Grep()

	if err != nil {
		log.Println("Unable to read the file")
		return
	}

	for _, x := range output {
		fmt.Println(x)
	}
}

func (a *App) Grep() ([]string, error) {

	output := make([]string, 0)
	reader := bufio.NewReader(a.source)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")

		if strings.Contains(line, a.substring) {
			output = append(output, line)
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return []string{}, err
		}

	}

	return output, nil
}
