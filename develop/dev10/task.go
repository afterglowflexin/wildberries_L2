package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

/*
=== Утилита telnet ===
Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123
Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

const (
	network = "tcp"
)

func main() {

	var timeoutDur string

	flag.StringVar(&timeoutDur, "timeout", "10s", "Timeout to connect")
	flag.Parse()

	timeout, err := GetDuration(timeoutDur)

	if err != nil {
		log.Panic(err)
	}

	host, port := os.Args[len(os.Args)-2], os.Args[len(os.Args)-1]
	address := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.DialTimeout(network, address, timeout)

	if err != nil {
		log.Panicln(err)
	}

	defer conn.Close()

	go func() {
		for {
			io.Copy(conn, os.Stdin)
		}
	}()

	go func() {
		for {
			io.Copy(os.Stdout, conn)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Kill, os.Interrupt)

	<-sigChan
}

func GetDuration(stringDur string) (time.Duration, error) {

	timeDur, err := time.ParseDuration(stringDur)

	if err != nil {
		return 0, err
	}

	timeout := timeDur * time.Second

	return timeout, nil
}
