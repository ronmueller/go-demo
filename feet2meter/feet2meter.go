package main

import (
	"fmt"
)

const (
	factor float64 = 0.3048
)

func main(){
	fmt.Print("Bitte ft eingeben: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * factor
	fmt.Println(output)
}
