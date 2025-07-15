package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}

	if true {
		i := 1
		for i <= 10 {
			fmt.Println(i)
			i = i + 1
		}
	}

	var stop int = 18
	for 1 == 1 {
		stop++
		if stop == 20 {
			break
		}
	}

	var sum = 0
	var str string = "t1e2s3t4s5t6r7i8n9g"

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	for _, value := range str {
		if value < 0 {
			continue
		}
		fmt.Println("-", value)
		sum += int(value)
	}
	fmt.Println("----", sum)

	a, b, c, d, e := 3, 4, 5, 6, 7

	if a == 3 {
		if b == 4 {
			if c == 5 {
				if d == 6 {
					if e == 7 {
						fmt.Println("Цикломатическая сложность: 5")
					}
				}
			}
		}
	}
}
