// Напишите функцию, которая принимает строку, разбивает её на слова и возвращает слайс слов, отсортированных по длине (от коротких к длинным)
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	a := "Hello lets go study"
	str := strArr(a)
	fmt.Println(str)
}

func strArr(s string) []string {
	words := strings.Fields(s)

	sort.Slice(words, func(i int, j int) bool {
		return len(words[i]) < len(words[j])
	})
	return words
}
