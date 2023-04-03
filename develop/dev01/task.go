/*
=== Базовая задача ===
Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.
Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

const (
	ntpPool = "0.beevik-ntp.pool.ntp.org"
)

func main() {

	l := log.New(os.Stderr, "", 0)

	response, err := ntp.Query(ntpPool)

	if err != nil {
		l.Println(err)
		os.Exit(10)
	}

	time := time.Now().Add(response.ClockOffset)

	fmt.Println(time.Format("2006-01-02 15:04:05.000"))
}
