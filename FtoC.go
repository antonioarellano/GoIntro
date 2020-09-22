package main

import "fmt"

func main() {
	var f float32
	fmt.Println("Conversor Fahrenheit a Celcius ")
	fmt.Println("Capture los grados Fahrenheit")
	fmt.Scan(&f)
	celcius := (f - 32) * 5 / 9
	fmt.Println("Grados celsius: ", celcius)
}
