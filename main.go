package main

import (
	"fmt"
)

func main() {
	arrayAndSlice()
}

//1 Array,slice
func arrayAndSlice() {
	// slice
	slice1 := []int{1, 2, 4, 4}
	slice2 := append(slice1, 5, 6, 7)
	fmt.Println(slice1, slice2)

	slice3 := []int{4, 5, 2, 4, 1}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
}
