package main

import "fmt"

func main() {
	Cascade()
	fmt.Println("Работа программы успешно завершена!")
	Cascade()
	fmt.Println("Работа программы успешно завершена во второй раз!")
}

func Cascade() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered from panic: %v\n", r)
		}
	}()
	panicFunc()
	fmt.Println("Работа функции успешно завершена!")
}

func panicFunc() {
	fmt.Println("Момент до паники")
	if true {
		panic("Управляемая паника")
	}
	fmt.Println("Функция с паникой без паники!")
}
