package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

/*
=== Утилита wget ===
Реализовать утилиту wget с возможностью скачивать сайты целиком
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//site := "https://habr.com/ru/news/t/726730/"

	if len(os.Args) < 2 {
		logger.Println("Не указан адрес сайта")
		return
	}

	site := os.Args[1]

	resp, err := http.Get(site)

	if err != nil {
		logger.Println("Не удается сделать запрос")
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		logger.Println("Не удается прочитать данные")
		return
	}

	os.WriteFile("index.html", body, 0666)

	if err != nil {
		logger.Println("Не удается сохранить данные в файл")
		return
	}
}
