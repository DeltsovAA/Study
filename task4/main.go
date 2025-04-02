// Напишите функцию, которая принимает слайс строк и возвращает новый слайс без повторяющихся элементов.
package main

import "fmt"

func main() {
	arr := []string{"a", "f", "c", "d", "c", "f"}
	fmt.Println(newArr(arr))
}

func newArr(a []string) []string {
	var result []string
	for _, v := range a {
		uniq := true
		for _, vv := range result {
			if vv == v {
				uniq = false
				break
			}
		}
		if uniq {
			result = append(result, v)
		}
	}
	return result
}
