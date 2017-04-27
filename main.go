package main

import (
	"fmt"
)

func main() {
	//x := 5
	//mapExample()
	//closureExample()
	//fmt.Println(fractional(x))
	//deferExample()
	//pointerExample()

	//fmt.Println(half(4))
	//fmt.Println(variadic(9, 12, 11, 42, 5, 3, 4, 5))

	// Closure
	/*
		add := func(x, y int) int {
			return x + y
		}
		fmt.Println(add(2, 3))
	*/

	var x *int
	var y int
	y = 0
	x = &y

	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(y)

}

func zero(xPtr *int) {
	*xPtr = 0
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

//2 Map
func mapExample() {

	mapexample := make(map[string]int)
	mapexample["key"] = 10
	mapexample["key1"] = 11
	fmt.Println(mapexample["key"])

	// delete
	delete(mapexample, "key")
	fmt.Println(mapexample)

	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	fmt.Println(elements)

	elements2 := map[string]map[string]string{
		"H": map[string]string{"name": "Hydrogen", "state": "gas"},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{"name": "Lithium", "state": "solid"},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
	}
	fmt.Println(elements2)

	if el, ok := elements2["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	if el, ok := elements2["kk"]; !ok {
		fmt.Println("null cmnr", el, ok)
	}
}

// 3 Closure, Recursive
func closureExample() {
	x := 0
	increament := func() int {
		x++
		return x
	}
	fmt.Println(increament())
	fmt.Println(increament())
}

func fractional(x int) int {
	if x == 0 {
		return 1
	}
	return x * fractional(x-1)
}

// 4 defer, panic, recover
func deferExample() {
	defer second()
	first()
}

func first() {
	fmt.Println("First")
}

func second() {
	fmt.Println("Second")
}

// 5 Pointer
func pointerExample() {
	var x *int // declare pointer x is a int pointer
	var y int  // declare variable y is a int number
	y = 0      // assign y value = 0
	x = &y     // assign pointer x is

	fmt.Println(x)  // gia tri cua x la dia chi cua y
	fmt.Println(&y) // dia chi cua y
	fmt.Println(*x) // lay gia tri cua y = 0
	fmt.Println(y)  // gia tri cua y = 0

	*x = 1 // gan gia tri con tro  x= 1

	fmt.Println(*x) // gia tri cua x = 1
	fmt.Println(y)  // gia tri cua y = 1
}

func half(x int) (float32, bool) {
	if x%2 == 0 {
		return float32(x / 2), true
	}
	return float32(x / 2), false
}

func variadic(args ...int) int {
	min := args[0]
	for _, value := range args {
		if value < min {
			min = value
		}
	}
	return min
}

func deferWithPanic() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

func square(x *float64) {
	*x = *x * *x
}

func swap(p1, p2 *int) {
	// temp := new(int)
	// temp = &p1
}
