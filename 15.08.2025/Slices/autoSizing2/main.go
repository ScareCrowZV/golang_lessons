package main

import "fmt"

func main() {
	x := make([]float64, 1024)

	fmt.Println(x, cap(x), len(x))

	x = append(x, 1)

	fmt.Println(x, cap(x), len(x))

	x = append(x, 2)

	fmt.Println(x, cap(x), len(x))
}
