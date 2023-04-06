package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func main() {

	l := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	var cmd string

	for {
		fmt.Scan(&cmd)

		switch cmd {
		case "cd":

			var dir string
			fmt.Scan(&dir)

			if err := os.Chdir(dir); err != nil {
				l.Printf("Error while changing dir:%s\n", err)
				return
			}

			l.Printf("Changed dir to %s\n", dir)

		case "pwd":

			dir, err := os.Getwd()

			if err != nil {
				l.Printf("Error getting a rooted path: %s\n", err)
				continue
			}

			log.Println(dir)

		case "echo":

			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				log.Printf("%s", scanner.Text())
			}

		case "kill":

			var pidStr string
			fmt.Scan(&pidStr)

			pid, err := strconv.Atoi(pidStr)

			if err != nil {
				l.Println("Wrong pid")
				continue
			}

			process, err := os.FindProcess(pid)

			if err != nil {
				l.Println("No such process")
				continue
			}

			if err = process.Kill(); err != nil {
				l.Println("Can't kill the process")
				continue
			}

		case "ps":

			processes := v3.process.Processes()
			l.Println(os.Getpid())

		default:
			l.Println("No such command")
		}

	}
}
