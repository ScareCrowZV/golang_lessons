package main

import "fmt"

func main() {

	var i int
	fmt.Println("Введите номер дня недели")
	fmt.Scan(&i)

	if i == 1 {
		fmt.Println("Понедельник")
	} else if i == 2 {
		fmt.Println("Вторник")
	} else if i == 3 {
		fmt.Println("Среда")
	} else if i == 4 {
		fmt.Println("Четверг")
	} else if i == 5 {
		fmt.Println("Пятница")
	} else if i == 6 {
		fmt.Println("Суббота")
	} else if i == 7 {
		fmt.Println("Воскресение")
	}

	switch i {
	case 1:
		fmt.Println("Понедельник")
	case 2:
		fmt.Println("Вторник")
	case 3:
		fmt.Println("Среда")
	case 4:
		fmt.Println("Четверг")
	case 5:
		fmt.Println("Пятница")
	case 6:
		fmt.Println("Суббота")
		fmt.Println("Это суббота")
	case 7:
		fmt.Println("Воскресение")
	default:
		fmt.Println("Введён номер дня недели из другой вселенной")
	}

	a := 3
	b := 12

	switch c := 9; {
	case a-c == b:
		println("a minus c equals b")
	case b%a == 0:
		println("b is divisible by a")
	case a%b == 2:
		break
	default:
	}

	d := 5
	switch vr := 1; d { //vr доступна только внутри блока switch
	case 5:
		fmt.Println(vr) // выведет 1
	default:
	}
	fmt.Println(d) // выведет 5
}
