package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

/*
=== Задача на распаковку ===
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.
Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fmt.Println(Unpack("a4bc2d5e"))
	fmt.Println(Unpack("abcd"))
	fmt.Println(Unpack("45"))
	fmt.Println(Unpack(""))
}

func Unpack(s string) (string, error) {
	var builder strings.Builder

	var prevRune rune

	for i := 0; i < utf8.RuneCount([]byte(s)); i++ {
		curRune := []rune(s)[i]

		if unicode.IsNumber(curRune) {

			if !(unicode.IsLetter(prevRune)) {
				return "", errors.New("Invalid string")
			}

			n, err := strconv.Atoi(string(curRune))
			if err != nil {
				return "", err
			}

			for j := 0; j < n-1; j++ {
				builder.WriteRune(prevRune)
			}

		} else if unicode.IsLetter(curRune) {

			builder.WriteRune(curRune)
			prevRune = curRune

		}
	}

	return builder.String(), nil
}
