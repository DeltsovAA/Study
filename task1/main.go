// 1. Написать функцию, которая использовует указатели, чтобы поменять местами два числа без третьей переменной.
package main

import "fmt"

func main() {
	a, b := 5, 10
	swapPtr(&a, &b)
	fmt.Println(a, b)
}

func swapPtr(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}
