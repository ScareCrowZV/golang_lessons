package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Это будет за потом")
	defer fmt.Println("Это будет потом")
	defer fmt.Println("Это будет перед потом")

	fmt.Println("Это будет сначала")

	i := 4
	defer func(n int) {
		fmt.Println("Вы загадали значение:", n)
	}(i)
	i = 3
	fmt.Println(i)

	f, err := os.Open("Такого файла не существует") // открыть файл
	if err != nil {
		fmt.Println("error opening file: %v", err)
		return
	}
	defer f.Close() // отложенное закрытие файла

}
