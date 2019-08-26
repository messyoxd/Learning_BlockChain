package main

import "fmt"

func main() {
	var array1 [5]int
	for i := 0; i < 5; i++ {
		array1[i] = i
	}
	array2 := [5]int{1, 2, 3, 4, 5}
	for _, value := range array2 {
		fmt.Println(value)
	}
	slice := array2[2:5]
	fmt.Println(slice)

	slice2 := make([]int, 5, 6)
	copy(slice2, slice)

	fmt.Println(slice2)
	/* append returns an concatenated array */
	fmt.Println(append(slice2, 6, 7, 8, 9))
}
