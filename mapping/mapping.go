package main 

import (
	"fmt"
)

func main() {
	
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
	
	if name, ok:= elements["Ne"]; ok {
		fmt.Println(name, ok)
	}
	
	elements_int := map[int]string{
		1: "Hydrogen",
		2: "Helium",
		10: "Roman will haben!",
	}
	
	if name_int, ok_int:= elements_int[10]; ok_int {
		fmt.Println(name_int, ok_int)
	}
}

