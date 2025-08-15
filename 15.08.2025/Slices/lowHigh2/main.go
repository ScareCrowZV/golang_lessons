package main

import "fmt"

func main() {

	// var arr = [5]float64{1, 2, 3, 4, 5}

	// sl := arr[:]
	// sl2 := arr
	// fmt.Println(arr, sl, sl2)

	// sl[4] = 100
	// fmt.Println(arr, sl, sl2)

	var sl3 []int64 = make([]int64, 1, 5)
	var sl4 []int64 = make([]int64, 1, 5)
	sl3 = append(sl3, 1)
	sl3 = append(sl3, 2)

	fmt.Println(sl3, len(sl3), cap(sl3))

	sl4 = append(sl3, 1)

	fmt.Println(sl3, sl4, append(sl3, 1))

}
