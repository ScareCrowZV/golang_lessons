package main

import (
	"fmt"
)

func main() {

	var input int
	fmt.Println("Введите число")
	fmt.Scan(&input)

	if input%2 == 0 {
		fmt.Println("input is divisible by 2")
	} else if input%3 == 0 {
		fmt.Println("input is divisible by 3")
	} else if input%5 == 0 {
		fmt.Println("input is divisible by 5")
	} else {
		fmt.Println("input probably is divisible by something else")
	}

	// var num int = 2
	// fmt.Println(num)
	// if true {
	// 	num := 1
	// 	if num == 1 {
	// 		fmt.Println(num)
	// 	}
	// }

	// fmt.Println(num)

}
