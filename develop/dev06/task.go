package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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

type Config struct {
	fields    string
	delimiter string
	separated bool
}

type App struct {
	config  Config
	columns []int
}

func NewApp() *App {
	return &App{config: Config{}}
}

func (a *App) Init() {
	var columnsStr, delimiterStr string

	flag.StringVar(&columnsStr, "f", "", "Choose fields to print")
	flag.StringVar(&delimiterStr, "d", "\t", "Delimiter")
	//flag.BoolVar(&separated, "f", , )

	flag.Parse()

	if columnsStr != "" {

		columns := strings.Split(columnsStr, ",")

		for _, x := range columns {

			number, err := strconv.Atoi(x)

			if err != nil {
				log.Printf("%s is not a number\n", x)
			} else {
				a.columns = append(a.columns, number)
			}
		}

		sort.Ints(a.columns)

	}
}

func (a *App) Scan() (string, error) {

	builder := strings.Builder{}

	lines, err := a.Reader(os.Stdin)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	for _, x := range lines {

		build := strings.Builder{}

		cols := strings.Split(x, "\t")

		for _, y := range a.columns {
			if y <= len(cols) {
				build.WriteString(cols[y-1])
			}
		}

		builder.WriteString(fmt.Sprintf("%s\n", build.String()))
	}

	return builder.String(), nil
}

func (a *App) Reader(rd io.Reader) ([]string, error) {

	lines := make([]string, 0)
	reader := bufio.NewReader(rd)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		lines = append(lines, line)

		if err != io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}

	return lines, nil
}

func main() {
	app := NewApp()
	app.Init()
	output, err := app.Scan()

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(output)
}
