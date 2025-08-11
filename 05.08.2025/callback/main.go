package main

import "fmt"

func main() {

	checkHigh := func(a, b int) bool {
		if a > b {
			return true
		} else {
			return false
		}
	}

	fmt.Println(whoIsHigher(6, 8, 7, checkHigh))

}

// Объявляем входящий параметр функции whoIsHigher сигнатурой функции, которую в неё можно передать.
// Выше в функции main() объявляем функцию как переменную checkHigh и передаём её
func whoIsHigher(a, b, c int, checkHigher func(a, b int) bool) int {
	var res int

	if checkHigher(a, b) {
		res = a
	} else {
		res = b
	}

	if b > res {
		if checkHigher(b, c) {
			res = b
		} else {
			res = c
		}
	}

	if c > res {
		if checkHigher(c, a) {
			res = c
		} else {
			res = a
		}
	}

	return res
}
