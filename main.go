package main

import (
	"fmt"
)

func main() {
	mapExample()
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
