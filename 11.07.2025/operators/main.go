package main

import "fmt"

func main() {

	number := 4
	number = 4 + 5
	fmt.Println(number)

	var assigne_var_a, assigne_var_b, assigne_var_c int
	assigne_var_a, assigne_var_b, assigne_var_c = 1, 2, 3
	fmt.Println(assigne_var_a, assigne_var_b, assigne_var_c)

	var assigne_var_a_1, assigne_var_b_2, assigne_var_c_3 int
	assigne_var_a_1, assigne_var_b_2, _ = 1, 2, 3
	fmt.Println(assigne_var_a_1, assigne_var_b_2, assigne_var_c_3)

	operand_test_1 := 4
	operand_test_2 := 6
	operand_test_result_1 := operand_test_1 + operand_test_2
	fmt.Println(operand_test_result_1)

	operand_test_3 := 4.5
	operand_test_4 := 5.5
	operand_test_result_2 := operand_test_3 + operand_test_4
	fmt.Println(operand_test_result_2)

	operand_test_5 := `Привет`
	operand_test_6 := ` всем!`

	operand_test_result_3 := operand_test_5 + operand_test_6
	fmt.Println(operand_test_result_3)

	const (
		operand_const_1 = iota * 2
		operand_const_2
		operand_const_3
	)
	fmt.Println(operand_const_1, operand_const_2, operand_const_3)

	operand_test_7 := 4
	operand_test_8 := 6
	operand_test_result_4 := operand_test_7 + operand_test_8
	fmt.Println(operand_test_result_4)

	operand_test_result_5 := operand_test_7 - operand_test_8
	fmt.Println(operand_test_result_5)

	operand_test_result_6 := operand_test_7 * operand_test_8
	fmt.Println(operand_test_result_6)

	operand_test_result_7 := operand_test_7 / operand_test_8
	fmt.Println(operand_test_result_7)

	operand_test_result_8 := operand_test_7 % operand_test_8
	fmt.Println(operand_test_result_8)

	operand_test_7++
	fmt.Println(operand_test_7)

	operand_test_8--
	fmt.Println(operand_test_8)

	var wrong_type_test_1 int8 = 0
	var wrong_type_test_2 int64 = 255
	wrong_type_result_1 := wrong_type_test_1 + int8(wrong_type_test_2)
	fmt.Println(wrong_type_result_1)

	// 2.718281828 преобразуется в тип float32
	explicit_type_a := float32(2.718281828)
	// 1.0 + 0.0i преобразуется в тип complex128
	explicit_type_b := complex128(1)
	// 0.5 преобразуется в тип float32
	explicit_type_c := float32(0.49999999)
	// 0.0 преобразуется в тип float64
	explicit_type_d := float64(-1e-1000)
	// "x" преобразуется в тип string
	explicit_type_e := string('x')
	// "♬" преобразуется в тип string
	explicit_type_f := string(0x266c)
	// Нельзя: Литерал 1.2 никак не получится представить, как int
	explicit_type_g := 3 //int(1.2)
	// Нельзя: 65.0 это вещественный числовой литерал
	explicit_type_h := 2 //string(65.0)
	// Переменная типа float32 преобразуется в float64
	explicit_type_i := float64(explicit_type_a)
	// 5 преобразуется в тип int64
	explicit_type_j := int64(5)
	// Переменная типа int64 преобразуется в int8
	explicit_type_k := int8(explicit_type_j)
	// Переменная типа float32 преобразуется в int
	explicit_type_l := int(explicit_type_a)
	// Нельзя: a - переменная вещественного числового типа
	explicit_type_m := 1 //string(explicit_type_a)
	// Код символа в виде литерала целого числа преобразуется в символьную
	// переменную
	explicit_type_o := rune(45)

	fmt.Println(explicit_type_a, explicit_type_b, explicit_type_c, explicit_type_d, explicit_type_e, explicit_type_f, explicit_type_g, explicit_type_h, explicit_type_i, explicit_type_j, explicit_type_k, explicit_type_l, explicit_type_m, string(explicit_type_o))

	overflow_a := 256
	overflow_b := uint8(overflow_a)

	fmt.Println(overflow_b)

	const explicit_2_a = 'H'
	const explicit_2_b = rune('H')

	fmt.Println(explicit_2_a, explicit_2_b)

	comparision_less_than := 3 < 4
	comparision_greater_than := 100 > 1
	comparision_a := 6
	comparision_not_equal := comparision_a != 3
	fmt.Println(comparision_less_than, comparision_greater_than, comparision_not_equal)

	lbool_is_true := 3 < 4 && 5 > 19
	lbool_and_is_this_true := 3+6 < 15
	fmt.Println(lbool_and_is_this_true, lbool_is_true)

	// test_comp_a := 15
	// test_comp_b := 14
	// test_comp_result := !test_comp_a < test_comp_b
	// fmt.Println(test_comp_result)

	const test_const_ex_a = iota + 1
	const (
		test_const_ex_b = 5
		test_const_ex_c = string('T')
		test_const_ex_d = iota + 5
	)

	fmt.Println(test_const_ex_a, test_const_ex_b, test_const_ex_c, test_const_ex_d)

	const ao = 3.14
	const ab = 5
	var bbb float64 = ao + ab
	// var ccc int = int(bbb)

	fmt.Println(ao, bbb)

}
