package main

import "fmt"

func main() {

	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var fl32 float32
	var fl64 float64
	var cplx64 complex64
	var cplx128 complex128
	var rn rune

	var str string

	var is bool

	fmt.Println("Значения по умолчанию")
	fmt.Println(i8, i16, i32, i64, fl32, fl64, cplx64, cplx128, rn, str, is)

	fmt.Println("Отрицательные значения")
	i8 = -128

	fmt.Println(i8)

	fmt.Println("Указание в двоичной системе")
	i8 = 0b01100100

	fmt.Println(i8)

	fmt.Println("Указание в восьмиричной системе")
	i8 = 0o010

	fmt.Println(i8)

	fmt.Println("Указание в шестнадцатиричной системе")
	i8 = 0xF

	fmt.Println(i8)

	fmt.Println("Инициализация на создании")
	varshort := 100

	fmt.Println(varshort)

	fmt.Println("Множественное объявление")
	var many_a, many_b, many_c int

	fmt.Println(many_a, many_b, many_c)

	fmt.Println("Множественное объявление в короткой форме")
	m_a, m_b, m_c := 0, 1, 2

	fmt.Println(m_a, m_b, m_c)

	fmt.Println("Способ инициализации нескольких переменных блоком")
	var (
		ma_a int
		ma_b int32
		ma_c uint64
	)

	fmt.Println(ma_a, ma_b, ma_c)

	fmt.Println("Float32")
	var realNumber float32

	fmt.Println(realNumber)

	fmt.Println("Выведение типов в int")
	realNumberShort := 100 // Это будет int

	fmt.Println(realNumberShort)

	fmt.Println("Выведение типов во float")
	realRealNumberShort := 100.0 // это будет float

	fmt.Println(realRealNumberShort)

	fmt.Println("Использование экспоненты и краткой записи дробей")
	realHexExp := 0x2.12p+10
	realHexExp2 := 0x.8p-0
	realDecExp := 6.67428e-11
	realDecExp2 := 1e6
	realDec := .25

	fmt.Println(realHexExp, realHexExp2, realDecExp, realDecExp2, realDec)

	fmt.Println("Текст и руны")
	var symbol rune
	symbol2 := 'Я'
	var symbol3 rune = '\u2626'

	fmt.Println(symbol, string(symbol2), string(symbol3))
	fmt.Printf("%#U", symbol3)
	fmt.Println("Многострочный текст")
	var stringU string = `lopata
	100500
	master`

	fmt.Println("Строка записанная с помощью escape символов и asii кодов")
	var stringUU string = "\nhi! \"\\ \077"
	fmt.Println(stringU, stringUU)
	fmt.Println(stringUU[5])

	fmt.Println("Тип BOOL")
	t := true
	f := false
	fmt.Println(t)
	fmt.Println(f)

	fmt.Println("Константы")
	const pi = 3.14
	const piT float64 = 3.14
	const piPlusOne = piT + 1
	const (
		constA int = 5
		constB     = 4e1
	)

	fmt.Println(pi, piT, piPlusOne, constA, constB)

	fmt.Println("Применение iota")
	const (
		c_iota_0 = iota
		c_iota_1 = iota
		c_iota_2 = iota
	)

	fmt.Println(c_iota_0, c_iota_1, c_iota_2)

	fmt.Println("Вычисление iota")
	const (
		c_iota_3 = 50 + iota
		c_iota_4 = c_iota_3 + iota
		c_iota_5 = 3
		c_iota_6 = 50 + iota
	)

	fmt.Println(c_iota_3, c_iota_4, c_iota_5, c_iota_6)

	fmt.Println("Ещё вариант вычисления iota")
	const (
		c_iota_7         = iota * 42
		c_iota_8 float64 = iota * 42
		c_iota_9         = iota * 42
	)

	fmt.Println(c_iota_7, c_iota_8, c_iota_9)

	fmt.Println("Пропуск iota")
	const (
		c_iota_10 = iota
		_         = iota
		c_iota_11 = iota
	)

	fmt.Println(c_iota_10, c_iota_11)

	const (
		c_iota_12 = iota
		_         = iota
		c_iota_13 = iota
	)

	fmt.Println(c_iota_12, c_iota_13)

	nr := "\U000065e5"
	n2 := nr + "123"
	fmt.Println(n2)
	// n := "\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"

}

// 	Первая форма — представление в виде шестнадцатеричной формы числа от 0 до 255 в формате:
//  '\x<число_в шестнадцатеричной_форме>'

// Второй формат также позволяет записать код только в диапазоне от 0 до 255, но в восьмеричной форме:
//  '\<число_в шестнадцатеричной_форме>'

// Третий формат позволяет записать код символа в диапазоне от 0 до 65535, число записывается в шестнадцатеричной форме:
//  '\u<число_в шестнадцатеричной_форме>'

// Четвёртый формат позволяет записать код символа в диапазоне от 0 до 0x10FFFF, число записывается в шестнадцатеричной форме:
//  '\U<число_в шестнадцатеричной_форме>'

// "日本語"                                 // UTF-8 input text
// `日本語`                                 // UTF-8 input text as a raw literal
// "\u65e5\u672c\u8a9e"                    // the explicit Unicode code points
// "\U000065e5\U0000672c\U00008a9e"        // the explicit Unicode code points
// "\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // the explicit UTF-8 bytes
// }

// uint8       the set of all unsigned  8-bit integers (0 to 255)
// uint16      the set of all unsigned 16-bit integers (0 to 65535)
// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

// int8        the set of all signed  8-bit integers (-128 to 127)
// int16       the set of all signed 16-bit integers (-32768 to 32767)
// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

// float32     the set of all IEEE 754 32-bit floating-point numbers
// float64     the set of all IEEE 754 64-bit floating-point numbers

// complex64   the set of all complex numbers with float32 real and imaginary parts
// complex128  the set of all complex numbers with float64 real and imaginary parts

// byte        alias for uint8
// rune        alias for int32
