package main

import (
	"fmt"
	"sort"
)

/*
=== Поиск анаграмм по словарю ===
Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.
Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	words := []string{"пятак", "Листок", "стол", "слиток", "слиток", "кошмар", "столик", "пятка"}

	fmt.Println(SearchAnagram(&words))
}

func SearchAnagram(words *[]string) map[string][]string {

	result := make(map[string][]string)

	helper := make(map[string][]int)

	for i, x := range *words {

		//x = strings.ToLower(x)

		slice := []rune(x)

		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})

		str := string(slice)

		helper[str] = append(helper[str], i)
	}

	for _, x := range helper {

		if len(x) > 1 {

			for _, y := range x {

				first := (*words)[x[0]]

				result[first] = append(result[first], (*words)[y])
			}
		}
	}

	for _, x := range result {
		sort.Strings(x)
	}

	return result
}
