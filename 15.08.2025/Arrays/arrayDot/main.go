package main

import "fmt"

func main() {
	intArr := [...]int{12, 26, 53, 55, 65, 63}
	fmt.Println(len(intArr))
	fmt.Println(intArr[2])

	// var intArr2 [10]int

	for i := 0; i < 7; i++ {
		intArr[i] = i + 1
	}
}
