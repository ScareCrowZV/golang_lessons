package main

import "fmt"

func main() {
	fmt.Println(manyParameters(1, 1, "Ivan", "Petya", "Vasya"))
}

func manyParameters(a int, b int, names ...string) string {

	sum := a + b
	var nameString string
	for _, val := range names {
		nameString += ", " + val
	}
	return "Hello, World! sum is " + string(sum) + ". Hi" + nameString
}
