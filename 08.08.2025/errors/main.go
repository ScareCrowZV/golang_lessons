package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	// Пример 1
	os.Stdout.Close()

	// Вариант с использованием переменной внутри if
	if nb, err := fmt.Println("test"); err != nil {
		fmt.Println(err, nb)
		log.Fatal(err)
	}

	// Вариант с получением ошибки и последующей проверкой
	nb, err := fmt.Println("test")
	if err != nil {
		log.Fatal(err)
		fmt.Println(err, nb)
	}

	// // Пример 2
	// a, err := custom_error(1, 2)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(a)

	// // Пример 3
	// a, err := custom_error_format(127, 127)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(a)
	// var zero, digit int = 0, 1

	// fmt.Println(digit / zero)

}

func custom_error(a, b int8) (int8, error) {

	if a != b {
		return 0, errors.New("this error from custom_error(): a != b")
	}

	return a + b, nil

}

func custom_error_format(a, b int8) (int8, error) {
	_, err := custom_error_format_inner(a, b)

	if a != b {
		additional_info := "this is very bad"
		return 0, fmt.Errorf("%s:a != b / %w", additional_info, err)
	}

	return a + b, nil
}

func custom_error_format_inner(a, b int8) (int8, error) {
	if a != b {
		return 0, fmt.Errorf("Всё таки a не равно b")
	}

	return 0, nil
}
