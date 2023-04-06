package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

/*
=== Утилита sort ===
Отсортировать строки (man sort)
Основное
Поддержать ключи
-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
Дополнительное
Поддержать ключи
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// use case: go run task.go args source

type Config struct {
	k int
	n bool
	r bool
	u bool
}

type App struct {
	source *os.File
	config Config
}

func NewApp() *App {
	return &App{config: Config{}}
}

func (a *App) Init() {
	flag.IntVar(&a.config.k, "k", 0, "Column for sort")
	flag.BoolVar(&a.config.n, "n", false, "Sort by numbers")
	flag.BoolVar(&a.config.r, "r", false, "Reverse order")
	flag.BoolVar(&a.config.u, "u", false, "Skip repeating lines")

	flag.Parse()

	source := os.Args[len(os.Args)-1]

	file, err := os.Open(source)

	if err != nil {
		log.Println("Unable to open the file")
		os.Exit(1)
	}

	a.source = file
}

func (a *App) Scan() ([]string, error) {
	data, err := a.Read()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	split := strings.Split(data, "\n")

	// обратный порядок
	if a.config.r {
		sort.Slice(split, func(i, j int) bool {
			return split[i] > split[j]
		})
	} else {
		sort.Strings(split)
	}

	// удаление повторяющихся строк
	if a.config.u {
		for i := 1; i < len(split); i++ {
			if split[i] == split[i-1] {
				copy(split[i:], split[i+1:])
				split[len(split)-1] = ""
				split = split[:len(split)-1]

				i--
			}

		}
	}

	return split, nil
}

func (a *App) Read() (string, error) {

	builder := strings.Builder{}

	reader := bufio.NewReader(a.source)

	for {

		line, err := reader.ReadString('\n')
		builder.WriteString(line)

		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
	}

	return builder.String(), nil

}

func main() {

	app := NewApp()
	app.Init()
	sortedData, err := app.Scan()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for _, x := range sortedData {
		fmt.Println(x)
	}
}
