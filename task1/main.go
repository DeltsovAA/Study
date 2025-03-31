// 2. Написать функцию, которая находить сумму элементов массива из 5 чисел.
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(intPtr(arr))
}

func intPtr(summ []int) int {
	var numSum int
	for _, v := range summ {
		numSum += v
	}
	return numSum
}
