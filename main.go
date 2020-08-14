package main

import (
	"fmt"
	"math"
)

var global_name = "Jordan"

//this is a little wrapper func because math is standardised on float64
func floor(f float32) float32 { return float32(math.Floor(float64(f))) }

func main() {
	//shorthand, var and const declarations
	local_name := "Kadish"
	var age int8 = 100
	var height float32 = 2.4
	const isCoder = true
	email, location := "jordan@email.com", "Never Never Land"
	fmt.Println(global_name, local_name, floor(height), age, isCoder, email, location)
	fmt.Printf("%T %T %T %T %T %T %T\n", global_name, local_name, height, age, isCoder, email, location)
}

//23:40 https://www.youtube.com/watch?v=SqrbIlUwR0U
