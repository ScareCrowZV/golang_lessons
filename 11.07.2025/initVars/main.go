package main

import (
	"log"
)

const (
	myConst = "hello"
)

// Глобальные переменные
var (
	var1 string
	var2 int
	var3 bool
)

func main() {
	// Использование log
	log.Println(var1, var2, var3, myConst)

}
