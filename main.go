// 3. Написать функцию, которая удаляет элемент слайса по индексу.
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	ind := 2
	arr = arrDel(arr, ind)
	fmt.Println(arr)
}

func arrDel(a []int, i int) []int {
	a = append(a[:i], a[i+1:]...)
	return a
}
