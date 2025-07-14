package main

import (
	"fmt"
	englishMessage "hello_world_citilink/english"
	russianMessage "hello_world_citilink/russian"
)

func main() {

	fmt.Println(englishMessage.GetMessage())
	fmt.Println(russianMessage.GetMessage())
}
