package main

import (
	"recursionLesson/recursionExample"
)

func main() {
	/*
		1 -- верхний элемент стека
			2 + 1
				2 + 2 + 1
					2 + 2 + 2 + 1
						2 + 2 + 2 + 2 + 1 -- последний элемент стека

		1
			2 + 1 = 3
				2 + 3 = 5
					2 + 5 = 7
						2 + 7 = 9

	*/
	recursionExample.ArithmeticProgressionElementWithPrint(1, 2, 5)

}
