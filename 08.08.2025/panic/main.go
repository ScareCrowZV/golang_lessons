package main

import "fmt"

func main() {
	checkValue(1)
	checkValue(2)
	checkValue(5)
}

func checkValue(a int8) {
	switch a {
	case 1:
		fmt.Println(a)
		return
	case 2:
		fmt.Println(a)
		return
	case 3:
		fmt.Println(a)
		return
	case 4:
		fmt.Println(a)
		return
	default:
		panic("Ой! Что творится!")
	}

}
