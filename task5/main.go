// Напишите функцию, которая принимает строку, состоящую из слов, разделенных пробелами, и возвращает самое длинное слово. Если несколько слов имеют одинаковую максимальную длину, верните первое из них
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello World my name is Anthony"
	fmt.Println(strRe(str))
}

func strRe(s string) string {
	str := strings.Fields(s)
	if len(str) == 0 {
		return ""
	}
	long := str[0]
	for _, v := range str {
		if len(v) > len(str) {
			long = v
		}
	}
	return long
}
