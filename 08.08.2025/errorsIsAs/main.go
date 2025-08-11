package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// Создаём новую ошибку
	errIsToManyErrors := errors.New("too many errors")
	// Создаём ещё одну ошибку и оборачиваем через Errorf первую созданную
	err := fmt.Errorf("wrap: %w", errIsToManyErrors)
	// Проверяем наличие первой ошибки errIsToManyErrors в стеке ошибок err
	if errors.Is(err, errIsToManyErrors) {
		fmt.Println("Эта именно та ошибка про кучу ошибок")
	} else {
		fmt.Println("Нет это какая-то другая ошибка")
	}

	// совпадает ли ошибка по типу с указанной. Если есть, то заполняем её параметрами ошибки
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
}
