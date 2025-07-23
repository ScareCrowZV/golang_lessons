package main

import (
	"fmt"
	"strconv"
)

func main() {
	var text string
	fmt.Print("Введите число: ")
	_, err := fmt.Scanln(&text)

	if err != nil {
		fmt.Println(fmt.Sprintf("Не получилось прочитать строку: %s", err))
	}

	// Конвертация строки с помощью пакета strconv
	_, err2 := strconv.Atoi(text)
	if err2 != nil {
		fmt.Println(fmt.Sprintf("Не получилось сконвертировать строку в число: %s", err2))
	}

	if err == nil {
		fmt.Print(text)
	}
}
