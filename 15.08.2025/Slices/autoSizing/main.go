package main

import "fmt"

func main() {
	slice := make([]float64, 5)
	fmt.Println(slice, cap(slice), len(slice))
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 5)
	slice = append(slice, 6)
	fmt.Println(slice, cap(slice), len(slice))
}
