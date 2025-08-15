package main

import "fmt"

func main() {
	arr := [5]float64{1, 2, 3, 4, 5}

	// sl := arr[low:high]
	x := arr[0:5]
	fmt.Println(arr, cap(arr), len(arr))
	fmt.Println(x, cap(x), len(x))
}
